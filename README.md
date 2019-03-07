# daily-content

Application used to consume daily content from the internet.

## Technologies used

* [Golang](https://golang.org/)
* [Docker](https://www.docker.com/)

## Run

First, clone the repository:

```bash
git clone https://github.com/mjakobczyk/minhash.git
```
Second, run docker-compose to run application without managing dependencies:

```bash
docker-compose build
docker-compose up
```

Alternatively, build and run manually:

```go
go build *.go
go run *.go
```