.PHONY: build version dockerize clean daily-content

version=0.1.0
app=daily-context

build: daily-content

version:
	@echo $(version)

dockerize:
	docker build -t $(app):$(version) .

clean: 
	rm main-content

daily-content: 
	CGO_ENABLED=0 go build -v -o main -mod=vendor ./cmd/