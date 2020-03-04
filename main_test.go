package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReplaceHostname(t *testing.T) {
	value := "https://www.foo.com/es/newpage"
	got := replaceHostname(value, "www.foo.com","www.newhost.com")
	want := "https://www.newhost.com/es/newpage"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCheckHttpReponse(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, client")
		}))
	defer ts.Close()
	var userAgent string = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	got := checkHttp(ts.URL, userAgent)
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestEmpty (t *testing.T) {
	var value string = ""
	got := empty(value)
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestGetEnvDefaultValue(t *testing.T){
	got := getEnv("FOOENV","default")
	want := "default"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestGetEnv(t *testing.T){
	const testKey = "FOO"
	const testValue = "value"
	if err := os.Setenv(testKey, testValue); err != nil {
		t.Fatalf("Setenv(%q, %q) failed: %v", testKey, testValue, err)
	}
	got := getEnv("FOO","default")
	want := "value"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
