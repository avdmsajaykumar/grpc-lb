package main

import (
	"ajaykumar/grpc-lb/proto"
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
)

func main() {

	con, err := grpc.Dial("dns:///avdms.io:5000",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := proto.NewSampleClient(con)

	for i := 0; i < 5; i++ {
		res, err := client.Hi(context.Background(), &proto.WelcomeReq{Msg: "Hi"})
		if err != nil {
			print(err)
		}
		fmt.Println(res.GetMsg())
	}
}
