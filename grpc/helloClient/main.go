package main

import (
	"flag"
	"fmt"
	"time"

	grpclb "github.com/cheyuexian/go-excise/grpc/discovery"
	pb "github.com/cheyuexian/go-excise/grpc/helloworld"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	reg = flag.String("reg", "http://139.196.166.222:8080", "register etcd address")
)

func main() {
	flag.Parse()
	fmt.Println("serv", *serv)
	r := grpclb.NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		panic(err)
	}
	fmt.Println("conn...")

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
		} else {
			fmt.Println(err)
		}
	}
}
