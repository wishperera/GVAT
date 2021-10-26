#!/usr/bin/env bash
PROJECT_NAME="${PWD##*/}"

GOOS=$(go env GOHOSTOS)
GOARCH=$(go env GOHOSTARCH)
BUILD_NAME=$PROJECT_NAME-$GOOS-$GOARCH
if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi
go build -v -o "${BUILD_NAME}"
./"${BUILD_NAME}"