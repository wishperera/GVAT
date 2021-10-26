# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /app/gvat

FROM alpine

WORKDIR /app

COPY --from=builder /app/gvat /app/gvat
EXPOSE 8080

CMD [ "/app/gvat" ]