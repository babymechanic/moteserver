package server

import (
	"fmt"
	"github.com/babymechanic/moteserver/components"
	"github.com/golang/glog"
	"net"
	"net/http"
	"net/rpc"
)

type MoteServer struct {
	port int
}

func New(port int) *MoteServer {
	return &MoteServer{port}
}

func (server *MoteServer) Start() {
	fmt.Println("starting server on port", server.port)
	rpc.Register(new(components.Display))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", server.port))
	if err != nil {
		fmt.Println("listen error: %s", err)
		glog.Fatalf("listen error: %s", err)
	}

	http.Serve(listen, nil)
	fmt.Println("started server")
}
