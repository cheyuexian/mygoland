package main

import (
	"context"
	"fmt"
	example "github.com/cheyuexian/go-excise/grpc-gateway/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:generate protoc -I/usr/local/include -I. \
//  -I$GOPATH/src \
//  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
//  --go_out=plugins=grpc:. \
//  path/to/your_service.proto

//p=G:/Downloads/protoc-3.11.4-win64/include
//pp=F:/gopath/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.13.0/third_party/googleapis
//--grpc-gateway_out=logtostderr=true:.
//protoc -I$p -I$pp -I./ helloworld.proto --go_out=plugins=grpc:./
// --grpc-gateway_out=
// protoc -I$p -I$pp -I./ helloworld.proto --grpc-gateway_out=logtostderr=true:.
const (
	port = ":50051"
)
type server struct{
	example.UnimplementedEchoServiceServer
}
func (s *server) Echo(ctx context.Context, req *example.StringMessage) (*example.StringMessage, error) {
	fmt.Println("echo server",req.GetValue())
	v := "res "+req.GetValue()
	return &example.StringMessage{Value:v}, nil

}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	example.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
