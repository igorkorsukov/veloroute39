package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"route"

	"github.com/constabulary/gb/testdata/src/c"
)

func main() {

	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("jsonrpc.Dial: ", err)
	}

	err = c.Call("Calculator.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
}

func routes(rpc *rpc.Client) ([]route.Route, error) {

	var rts []route.Route
	rpc.Call("RouteService.Routes")
}
