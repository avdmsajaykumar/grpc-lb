package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	pb "ajaykumar/grpc-lb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	pb.UnimplementedSampleServer
}

func (s *server) Hi(ctx context.Context, in *pb.WelcomeReq) (*pb.WelcomeRes, error) {
	println("welcome")
	return &pb.WelcomeRes{Msg: *port}, nil
}

var (
	port = flag.String("port", "10000", "The server port")
)

func main() {

	flag.Parse()

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 10 * time.Second,
		}),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		os.Exit(1)
	}
	grpcserver := grpc.NewServer(opts...)
	pb.RegisterSampleServer(grpcserver, &server{})
	fmt.Println(*port)
	_ = grpcserver.Serve(lis)

}
