package main

import (
	"google.golang.org/grpc"
	"log"
	pb "prototest"
	"golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial:%v", err)
	}
	defer conn.Close()
	client := pb.NewTestServiceClient(conn)

	testAB(client)
}

func testAB(client pb.TestServiceClient) {
	ret, err := client.AB(context.Background(), &pb.A{Data:"客户端"})
	if err != nil {
		log.Fatalf("%v.AB(_) = _, %v", client, err)
	}
	log.Println("返回:%v", ret.GetData())
}
