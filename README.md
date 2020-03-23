# Daily Content

[![GoDoc](https://godoc.org/github.com/mjakobczyk/daily-content?status.svg)](https://godoc.org/github.com/mjakobczyk/daily-content)
[![Go Report Card](https://goreportcard.com/badge/mjakobczyk/daily-content)](https://goreportcard.com/report/github.com/mjakobczyk/daily-content)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](https://github.com/mjakobczyk/daily-content/blob/master/LICENSE)

## Overview

Application used to consume daily content from the internet.

## Technologies used

* [Golang](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [Make](https://www.gnu.org/software/make/)

## Integration with services

* Powered by [NewsAPI.org](https://newsapi.org/)

## Run

First, clone this repository:
```bash
$ git clone https://github.com/mjakobczyk/daily-content.git
```

 Please check environment variables values provided in `values.env` file before running the application. If you want to integrate with `NewsAPI` service you have to provide your own API key - there is no default value.

Running application in stand-alone mode is not supported. Make use of suggested options:

### Docker Compose

`Docker Compose` is used for local development purposes.

```bash
$ cd deployments/docker-compose
$ docker-compose build
$ docker-compose up
```
