#!/usr/bin/env bash
pwd
if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi
docker build  -t "gvat" .
docker run -dp ${SERVER_PORT}:${SERVER_PORT} --env-file .env  --name gvat gvat