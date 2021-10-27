.DEFAULT:

all: build

build-and-run:
	sh cmd/build_and_run.sh

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o bin/gvat-darwin-amd64

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o bin/gvat-linux-amd64

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o bin/gvat-windows-amd64

build-binaries: build-darwin-amd64 build-linux-amd64 build-windows-amd64

clean-builds:
	rm -rf bin/*

test:
	go test -v ./...

lint:
	golangci-lint run -v

docker:
	sh cmd/docker_build_and_run.sh

mocks:
	sh cmd/generate_mocks.sh

configure-env:
	sh cmd/configure_env.sh

clean-docker:
	docker stop gvat
	docker rmi $(docker image ls | grep 'gvat')
