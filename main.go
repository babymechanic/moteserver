package main

import (
	"flag"
	"github.com/babymechanic/moteserver/server"
)

func main() {
	flag.Parse()
	server.New(8080).Start()
}
