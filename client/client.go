package main

import (
	"ajaykumar/grpc-lb/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	con, _ := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))

	client := proto.NewSampleClient(con)

	res, err := client.Hi(context.Background(), &proto.WelcomeReq{Msg: "Hi"})
	if err != nil {
		print(err)
	}
	fmt.Print(res.GetMsg())
}
