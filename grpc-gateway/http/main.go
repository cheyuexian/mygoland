package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"fmt"
	"google.golang.org/grpc/connectivity"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	example "github.com/cheyuexian/go-excise/grpc-gateway/pb"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:50051", "gRPC server endpoint")
)

type server struct{
	example.UnimplementedEchoServiceServer
}
func (s *server) Echo(ctx context.Context, req *example.StringMessage) (*example.StringMessage, error) {
	fmt.Println("echo server",req.GetValue())
	v := "res "+req.GetValue()
	return &example.StringMessage{Value:v}, nil

}
func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}
func swaggerServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}
func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := example.RegisterEchoServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	server := http.Server{
		Handler:mux,
	}
	//mux.Handle("/swagger/",runtime.Pattern{}, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//	fmt.Println("swagger")
	//	swaggerServer("g/code/go/github.com/cheyuexian/go-excise/grpc-gateway/pb")
	//})

//	mux.HandleFunc("/swagger/", swaggerServer("g/code/go/github.com/cheyuexian/go-excise/grpc-gateway/pb"))

	lis, err := net.Listen("tcp", ":8080")
	fmt.Println("http err ",err)
	return server.Serve(lis)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	//return http.ListenAndServe(*grpcServerEndpoint, mux)
}
func grpc_server(){

	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println("errgrpc",err)
	s := grpc.NewServer()
	example.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn,err := grpc.DialContext(ctx, *grpcServerEndpoint, grpc.WithInsecure())

	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			glog.Errorf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", swaggerServer("g:/code/go/github.com/cheyuexian/go-excise/grpc-gateway/pb"))
	mux.HandleFunc("/healthz", healthzServer(conn))
	muxx := runtime.NewServeMux()
	example.RegisterEchoServiceHandler(ctx,muxx,conn)



	mux.Handle("/", muxx)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	glog.Infof("Starting listening at %s", "8080")
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}
func main() {
	flag.Parse()
	defer glog.Flush()

	//fmt.Println(err)
	go grpc_server()

	Run()
	//if err := run(); err != nil {
	//	glog.Fatal(err)
	//}
}