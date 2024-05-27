package main

import (
	"context"
	"net"
	pb "project/go/grpc/protoc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServerServer
}

func (s *server) ServiceReply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	data := in.SomeData
	return &pb.Response{Reply: "this is message:" + data}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	newGrpcServer := grpc.NewServer()
	pb.RegisterMyServerServer(newGrpcServer, &server{})

	if err := newGrpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
