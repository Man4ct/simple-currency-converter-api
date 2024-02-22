build:
	docker build -t currency-converter-api .
run:
	docker run -p 8080:8080 currency-converter-api

.PHONY: build run
	