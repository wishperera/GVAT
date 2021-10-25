
.SILENT:

all: build

build:
	sh cmd/build.sh

test:
	go test -v ./...

lint:
	golangci-lint run -v

docker:
	sh cmd/docker_build.sh

mocks:
	sh cmd/generate_mocks.sh

