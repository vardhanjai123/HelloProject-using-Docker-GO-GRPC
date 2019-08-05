package main

import (
	"context"
	//"fmt"
	"time"
	//"os"
	//"strings"
	pb "proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewSplitServiceClient(conn)

	var s string = "Hello"
	//fmt.Scanln(&s)
	for i:=1;i<=100;i++ {
		req := &pb.Request{S: s}
		if _, err := client.Split(context.Background(), req); err == nil {
			
		}
		time.Sleep(5*time.Second)
	}

	
}