package main
//
//import (
//	"flag"
//	"fmt"
//	pb "github.com/cheyuexian/go-excise/grpc/helloworld"
//	"github.com/golang/groupcache/testpb"
//	"google.golang.org/grpc/balancer/roundrobin"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/peer"
//	"google.golang.org/grpc/resolver"
//	"google.golang.org/grpc/resolver/manual"
//	"google.golang.org/grpc/status"
//	"time"
//
//	grpclb "github.com/cheyuexian/go-excise/grpc/discovery"
//
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"log"
//	"strconv"
//)
//
//var (
//	serv = flag.String("service", "hello_service", "service name")
//	reg = flag.String("reg", "http:http://139.196.166.222:8080", "register etcd address")
//)
//
//func TestBackendsRoundRobin() {
//	r, cleanup := manual.GenerateAndRegisterManualResolver()
//	defer cleanup()
//
//	backendCount := 5
//
//
//
//	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
//	if err != nil {
//		log.Fatal("failed to dial: %v", err)
//	}
//	defer cc.Close()
//	testc:=pb.NewGreeterClient(cc)
//	// The first RPC should fail because there's no address.
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
//	defer cancel()
//	//resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
//	reqs := &pb.HelloRequest{Name:"che"}
//	if _, err := testc.SayHello(ctx, reqs); err == nil || status.Code(err) != codes.DeadlineExceeded {
//		log.Fatal("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
//	}
//
//	var resolvedAddrs []resolver.Address
//	for i := 0; i < backendCount; i++ {
//		resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: test.addresses[i]})
//	}
//
//	r.UpdateState(resolver.State{Addresses: resolvedAddrs})
//	var p peer.Peer
//	// Make sure connections to all servers are up.
//	for si := 0; si < backendCount; si++ {
//		var connected bool
//		for i := 0; i < 1000; i++ {
//			if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
//				log.Fatal("EmptyCall() = _, %v, want _, <nil>", err)
//			}
//			if p.Addr.String() == test.addresses[si] {
//				connected = true
//				break
//			}
//			time.Sleep(time.Millisecond)
//		}
//		if !connected {
//			log.Fatal("Connection to %v was not up after more than 1 second", test.addresses[si])
//		}
//	}
//
//	for i := 0; i < 3*backendCount; i++ {
//		if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
//			log.Fatal("EmptyCall() = _, %v, want _, <nil>", err)
//		}
//		if p.Addr.String() != test.addresses[i%backendCount] {
//			log.Fatal("Index %d: want peer %v, got peer %v", i, test.addresses[i%backendCount], p.Addr.String())
//		}
//	}
//}
//func main() {
//	flag.Parse()
//	r := grpclb.NewResolver(*serv)
//	b := grpc.RoundRobin(r)
//
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
//	if err != nil {
//		panic(err)
//	}
//
//	ticker := time.NewTicker(1 * time.Second)
//	for t := range ticker.C {
//		client := pb.NewGreeterClient(conn)
//		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
//		if err == nil {
//			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
//		}else{
//			fmt.Println("error ",err)
//		}
//	}
//}