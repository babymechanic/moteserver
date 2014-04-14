package main

import (
	"github.com/babymechanic/moteserver/server"
)

func main() {
	server.New(8080).Start()
}
