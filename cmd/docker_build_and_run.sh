#!/usr/bin/env bash
IMAGE_NAME="gvat"
if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi
docker build  -t "$IMAGE_NAME" .
docker image prune -f

if docker image ls -a "$IMAGE_NAME" | grep -Fq "$IMAGE_NAME" 1</dev/null; then
  docker rm "$IMAGE_NAME"
fi

docker run -dp ${SERVER_PORT}:${SERVER_PORT} --env-file .env  --name gvat gvat