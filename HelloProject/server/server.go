package main

import (
	"context"
	"net"
	"fmt"
	//"strings"
	pb "proto"
	"google.golang.org/grpc"
)
var count int64 = 1
type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterSplitServiceServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (server) Split(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	s := request.GetS()
	//fmt.Println(s)

	//split := strings.Split(s, "")
	//return &pb.Response{Items: split}, 
	fmt.Printf("Jaivardhan! %s %dst time",s,count)
	fmt.Println()
	count++
	return &pb.Response{}, nil
	
}