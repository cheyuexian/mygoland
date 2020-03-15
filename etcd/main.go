package  main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var  endpoints []string;
var dialTimeout time.Duration = 10*time.Second
var requestTimeout time.Duration = 10*time.Second

func ExampleKV_put() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleKV_get() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	// Output: foo : bar
}
func ExampleLease_grant() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	//defer cli.Close()

	// minimum lease TTL is 5-second
	var interval int64  = 15
	resp, err := cli.Grant(context.TODO(), interval)
	if err != nil {
		log.Fatal(err)
	}
	key := "foo"
//	after 5 seconds, the key 'foo' will be removed
	_, err = cli.Put(context.TODO(), key, "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cli",&cli,cli)
	var stopSignal = make(chan bool, 1)
	resp,_ = cli.Grant(context.TODO(),interval)
	fmt.Println("123")
	go func() {
		ticker := time.NewTicker(10*time.Second)
		for   {
			fmt.Println("tyestaaaaaaaaaaaaaaaaaaaaa")
			fmt.Println("cli",&cli,cli)
			//resp,_ := cli.Grant(context.TODO(),interval)
			resp, err := cli.Grant(context.TODO(), interval)
			fmt.Println("grant ",resp,err)
			resp1,err  := cli.Get(context.Background(),"foo")
			fmt.Println("aaaaa",err)
			for _, ev := range resp1.Kvs {
				fmt.Printf("kv %s : %s\n", ev.Key, ev.Value)
			}
			if err != nil{
			//	fmt.Println("grant cli ")
			}else{
				if _,err = cli.Put(context.TODO(),key,"bar",clientv3.WithLease(resp.ID)); err  != nil{
				//	fmt.Println("put cli",err)
				}
			}
			select{
			case <-stopSignal:
				return
			case <- ticker.C:

			}

		}
	}()

}
func main(){
	endpoints  = append(endpoints, "http://139.196.166.222:8080")


	//ExampleKV_put()
	//ExampleKV_get()
	ExampleLease_grant()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	s := <-ch
	log.Printf("receive signal '%v'", s)
	os.Exit(1)
	go func() {

	}()



}