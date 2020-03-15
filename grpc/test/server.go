package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	//"google.golang.org/grpc/internal/leakcheck"
	testpb "google.golang.org/grpc/test/grpc_testing"

)
type test1 struct {
	servers   []*grpc.Server
	addresses []string
}

type testServer1 struct {
	testpb.UnimplementedTestServiceServer
}
func main21() {

	ch := make(chan int)
	startTestServers1(1)

	<-ch


}
func startTestServers1(count int) (_ *test1, err error) {
	t := &test1{}

	defer func() {
		if err != nil {
			//t1.cleanup()
		}
	}()
	for i := 0; i < count; i++ {
		lis, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			return nil, fmt.Errorf("failed to listen %v", err)
		}

		s := grpc.NewServer()
		testpb.RegisterTestServiceServer(s, &testServer1{})
		t.servers = append(t.servers, s)
		t.addresses = append(t.addresses, lis.Addr().String())

		go func(s *grpc.Server, l net.Listener) {
			s.Serve(l)
		}(s, lis)
	}

	return t, nil
}
