package main

import (
	"golang.org/x/net/context"
	pb "prototest2"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"io"
	"github.com/golang/protobuf/proto"
)

type TestService struct {

}

func (this *TestService) AB(ctx context.Context, a *pb.A) (*pb.B, error) {
	return &pb.B{Data: proto.String("ret:"+a.GetData())}, nil
}

func (this *TestService) ABB(a *pb.A, stream pb.TestService_ABBServer) error {
	for i := 0; i < 3; i++ {
		if err := stream.Send(&pb.B{Data:proto.String(fmt.Sprintf("[%v]%v", i, a.GetData()))}); err != nil {
			return err
		}
	}
	return nil
}

func (this *TestService) AAB(stream pb.TestService_AABServer) error {
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.B{Data:proto.String(fmt.Sprintf("次数:%v", count))})
		}
		if err != nil {
			return err
		}
		count++
	}
	return nil
}

func (this *TestService) AABB(stream pb.TestService_AABBServer) error {
	for {
		a, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for i := 0; i < 3; i++ {
			if err := stream.Send(&pb.B{Data:proto.String(fmt.Sprintf("[%v]%v", i, a.GetData()))}); err != nil {
				return err
			}
		}
	}
	return nil
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
