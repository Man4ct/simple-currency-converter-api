build:
	docker-compose up --build
run:
	docker run -p 8080:8080 currency-converter-api

.PHONY: build run
	