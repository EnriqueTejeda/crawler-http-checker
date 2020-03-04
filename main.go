package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/joho/godotenv"
	"github.com/yterajima/go-sitemap"
)

type config struct {
	host            string
	hostReplace     string
	userAgent       string
	workerSize      string
	sitemapFilename string
}

var netClient = &http.Client{}

func init() {
	tr := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	netClient = &http.Client{Timeout: 60 * time.Second, Transport: tr}
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: not found .env file, maybe you run this in a container.")
	}
}

func main() {
	start := time.Now()
	config := config{
		host:            getEnv("HOST", ""),
		hostReplace:     getEnv("NEW_HOST", ""),
		userAgent:       getEnv("USER_AGENT", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		workerSize:      getEnv("WORKER_SIZE", "1"),
		sitemapFilename: getEnv("SITEMAP_FILENAME", "sitemap.xml"),
	}
	if empty(config.host) {
		log.Fatalln("Error hostname not found")
	}
	workerSize, err := strconv.Atoi(config.workerSize)
	if err != nil {
		log.Fatalln("Error parsing workerSize variable")
	}
	wp := workerpool.New(workerSize)
	smap, err := sitemap.Get(config.host+"/"+config.sitemapFilename, nil)
	if err != nil {
		log.Fatalln(err)
	}
	for _, URL := range smap.URL {
		URL := URL
		if !empty(config.hostReplace) {
			URL.Loc = replaceHostname(URL.Loc, config.host, config.hostReplace)
		}
		wp.Submit(func() {
			if checkHttp(URL.Loc, config.userAgent) {
				log.Printf("%s is fine.", URL.Loc)
			}
		})
	}
	wp.StopWait()
	elapsed := time.Since(start)
	log.Printf("%d urls scanned in: %v", len(smap.URL), elapsed)
}

func replaceHostname(s string, old string, new string) string {
	return strings.Replace(s, old, new, 1)
}

func checkHttp(urlRequest string, userAgent string) bool {
	var ok bool = true
	req, _ := http.NewRequest("GET", urlRequest, nil)
	req.Close = true
	req.Header.Add("User-Agent", userAgent)
	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ok = true
		log.Fatalf("Error in %q with http code : %v", urlRequest, resp.StatusCode)
	}
	return ok
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
