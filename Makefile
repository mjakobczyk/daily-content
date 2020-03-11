.PHONY: build version dockerize clean daily-content

version=0.1.0

build: daily-content

version:
	@echo $(version)

dockerize:
	docker build -t daily-content:$(version) .

clean: 
	rm daily-content

daily-content: 
	CGO_ENABLED=0 go build -o daily-content -mod=vendor *.go