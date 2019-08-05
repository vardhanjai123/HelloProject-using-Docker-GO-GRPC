FROM golang
ADD . /go/src/
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
WORKDIR /go/src/
