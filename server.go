package main

import (
	"fmt"
	"github.com/babymechanic/server/interfaces/components"
	"github.com/golang/glog"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	port := 8080
	fmt.Println("starting server on port", port)
	rpc.Register(new(components.Display))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Println("listen error: %s", err)
		glog.Fatalf("listen error: %s", err)
	}

	http.Serve(listen, nil)
	fmt.Println("started server")
}
