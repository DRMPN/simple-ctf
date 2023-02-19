FROM golang:1.13-alpine

WORKDIR /home/hometask/container
COPY container ./

CMD go run service.go & go run checker.go 172.26.13.11:7777 check