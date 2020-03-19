.PHONY: build version dockerize clean daily-content

version=0.1.0
app=daily-context

build: daily-content

version:
	@echo $(version)

dockerize:
	docker build -t $(app):$(version) .
