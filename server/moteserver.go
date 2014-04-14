package server

import (
	"fmt"
	_ "github.com/babymechanic/moteserver/components"
	"github.com/golang/glog"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type MoteServer struct {
	port int
}

func New(port int) *MoteServer {
	return &MoteServer{port}
}

func (server *MoteServer) Start() {
	fmt.Println("starting server on port", server.port)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", server.port))
	if err != nil {
		fmt.Println("listen error: %s", err)
		glog.Fatalf("listen error: %s", err)
		os.Exit(1)
	}
	http.Serve(listen, nil)
}
