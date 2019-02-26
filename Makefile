.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-all-user-func src/handlers/get-all-user-func/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-user-by-id-func src/handlers/get-user-by-id-func/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
