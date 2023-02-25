FROM golang:1.13-alpine

WORKDIR /home/hometask/container
COPY container ./

CMD go run service.go & go run checker.go host.docker.internal:7777 check