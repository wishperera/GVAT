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

#sudo sysctl net.ipv4.conf.all.forwarding=1
#sudo iptables -P FORWARD ACCEPT

docker run -dp ${SERVER_PORT}:${SERVER_PORT} --dns 8.8.8.8 --env-file .env --name gvat gvat