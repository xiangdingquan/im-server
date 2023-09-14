package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	pb "github.com/go-kratos/kratos/pkg/net/rpc/warden/internal/proto/testproto"
)

// usage: ./client -grpc.target=test.service=127.0.0.1:9000
func main() {
	log.Init(&log.Config{Stdout: true})
	flag.Parse()
	conn, err := warden.NewClient(nil).Dial(context.Background(), "direct://default/127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	cli := pb.NewGreeterClient(conn)
	normalCall(cli)
}

func normalCall(cli pb.GreeterClient) {
	reply, err := cli.SayHello(context.Background(), &pb.HelloRequest{Name: "tom", Age: 23})
	if err != nil {
		panic(err)
	}
	fmt.Println("get reply:", *reply)
}
