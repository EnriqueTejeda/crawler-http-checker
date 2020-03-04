![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/enriquetejeda/crawler-http-checker)
![Docker Pulls](https://img.shields.io/docker/pulls/etejeda/crawler-http-checker)
![GitHub top language](https://img.shields.io/github/languages/top/enriquetejeda/crawler-http-checker)
![GitHub last commit](https://img.shields.io/github/last-commit/enriquetejeda/crawler-http-checker)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Crawler Http Checker

Simple crawler for check the http code for all urls of your site, all made in Go inside a distroless container (only 11MB~)!

Visit also [docker hub repository](https://hub.docker.com/repository/docker/etejeda/crawler-http-checker).

## How works?

Basically is a simple app that parse your sitemap.xml and then make a request for each url in your site, is very usable for CI integration (for example test each url and terminate if detect any url with a bad http code).

Also you can configure the size of the workers pool for increase the number of parallel task and process all urls more quickly (default value is `1`)!.

## Requirements

* Docker Engine. :heart:

## Getting Started

### Docker :heart:

You only run this command in your terminal:

```
docker run \
-e 'HOST=https://www.enriquetejeda.com' \
etejeda/crawler-http-checker:latest
```

### Standalone

1. Rename the `.env.example` to `.env` and configure the values
2. Compile with the command `make build`
3. Run the command `make run`

### Continuous Integration
#### Jenkins
```
#!groovy
pipeline {
    agent { node { label 'master' } }
    options { skipDefaultCheckout true }
    environment {}
    stages {
        stage('Build'){
            steps {
                checkout scm
            }
        }
        stage('Test'){
            steps {
                echo 'Verify all urls..'
                docker.image('etejeda/crawler-http-checker:latest').run('-e HOST=https://www.enriquetejeda.com')
            }
        }
        stage('Deploy'){
            steps {
                echo 'deploy'
            }
        }
    }   
}
```

## Development

### Building the binary

I provided a makefile for do this job, only run this command:
```
make build 
```
### Building the container

I provided a makefile for do this job, only run this command:
```
make build-docker
```
### Environment Variables 

| Name  | Description  | Default | Required |
| -- | -- | -- | -- |
| HOST | The host for scan  | - | *yes* |
| NEW_HOST | If you require replace the url for other | - | *no* |
| USER_AGENT | User-Agent use for make each request | `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)` | *no* |
| WORKER_SIZE | The number of parallel task | `1` | *no* |
| SITEMAP_FILENAME | Name of the file for sitemap | `sitemap.xml` | *no* |

## How contribute? :rocket:

Please feel free to contribute to this project, please fork the repository and make a pull request!. :heart:

## Share the Love :heart:

Like this project? Please give it a ★ on [this GitHub](https://github.com/EnriqueTejeda/crawler-http-checker)! (it helps me a lot).

## License

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

See [LICENSE](LICENSE) for full details.

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

