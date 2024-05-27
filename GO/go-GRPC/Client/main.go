package main

import (
	"context"
	"fmt"
	pb "project/go/grpc/protoc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewMyServerClient(conn)

	ctc, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := c.ServiceReply(ctc, &pb.Request{SomeData: "myData"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
