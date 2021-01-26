.PHONY: clean build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main

clean:
	rm -rf ./bin

docker:
	docker-compose up --build -d

docker-stop:
	docker-compose down
