FROM golang:latest
MAINTAINER Tiny "xchliu.github.io"
WORKDIR /yocto
COPY build/yoctodb /yocto
COPY data /data
ENTRYPOINT  ["/yocto/yoctodb"]

#GOOS=linux go build -ldflags "-w" -o build/yoctodb  src/server/yoctod.go
#docker build -t yoctodb ./build/