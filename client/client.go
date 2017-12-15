package main

import (
	"google.golang.org/grpc"
	"log"
	pb "prototest2"
	"golang.org/x/net/context"
	"io"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial:%v", err)
	}
	defer conn.Close()
	client := pb.NewTestServiceClient(conn)

	testAB(client)
	testABB(client)
	testAAB(client)
	testAABB(client)
}

func testAB(client pb.TestServiceClient) {
	log.Println("testAB")
	ret, err := client.AB(context.Background(), &pb.A{Data:proto.String("客户端")})
	if err != nil {
		log.Fatalf("%v.AB(_) = _, %v", client, err)
	}
	log.Println("返回:" + ret.GetData())
}

func testABB(client pb.TestServiceClient) {
	log.Println("testABB")
	stream, err := client.ABB(context.Background(), &pb.A{Data:proto.String("调用ABB")})
	if err != nil {
		log.Fatalf("%v.ABB(_) = _,%v", client, err)
	}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ABB(_) = _,%v", client, err)
		}
		log.Println(b.GetData())
	}
}

func testAAB(client pb.TestServiceClient) {
	log.Println("testAAB")
	stream, err := client.AAB(context.Background())
	if err != nil {
		log.Fatalf("%v.AAB(_) = _,%v", client, err)
	}
	for i := 0; i < 3; i++ {
		err = stream.Send(&pb.A{Data:proto.String("调用AAB")})
		if err != nil {
			log.Fatalf("%v.AAB(_) = _,%v", client, err)
		}
	}
	b, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.AAB(_) = _,%v", client, err)
	}
	log.Println(b.GetData())
}

func testAABB(client pb.TestServiceClient) {
	log.Println("testAABB")
	stream, err := client.AABB(context.Background())
	if err != nil {
		log.Fatalf("%v.AABB(_) = _,%v", client, err)
	}

	waitc := make(chan struct{})
	go func(){
		for {
			b, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("%v.AABB(_) = _,%v", client, err)
			}
			log.Println(b.GetData())
		}
	}()

	for i := 0; i < 3; i++ {
		err := stream.Send(&pb.A{Data:proto.String(fmt.Sprintf("[%v]调用AABB", i))})
		if err != nil {
			log.Fatalf("%v.AABB(_) = _,%v", client, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

