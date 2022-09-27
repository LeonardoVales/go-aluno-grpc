FROM golang:1.18-alpine

WORKDIR /go/src

RUN apk update && apk add protobuf & \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest & \
    go install github.com/ktr0731/evans@latest
    

CMD ["tail","-f","/dev/null"]