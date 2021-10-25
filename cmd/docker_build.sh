#!/usr/bin/env bash
source cmd/configure.sh

docker build --build-arg BIN="${PROJECT_NAME}" -t "${PROJECT_NAME}:latest" .
