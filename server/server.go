package main

import (
	"golang.org/x/net/context"
	pb "prototest"
	"net"
	"log"
	"google.golang.org/grpc"
)

type TestService struct {

}

func (this *TestService) AB(ctx context.Context, a *pb.A) (*pb.B, error) {
	return &pb.B{Data: "ret:"+a.Data}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, &TestService{})
	grpcServer.Serve(lis)
}
