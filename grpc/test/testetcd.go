package main

import (
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"time"
	"context"
)

func main1(){

	_, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://254.0.0.1:12345"},
		DialTimeout: 2 * time.Second,
	})

	// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
	if err == context.DeadlineExceeded {
		// handle errors
	}

	// etcd clientv3 <= v3.2.9, grpc/grpc-go <= v1.2.1
	if err == grpc.ErrClientConnTimeout {
		// handle errors
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
}
