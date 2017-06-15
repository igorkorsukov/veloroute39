package main

import (
	"fmt"
	"infra"
	"infra/datastore"
	"net/http"
	"route"
)

func main() {

	fmt.Println("Hello world, I am veloroute39")

	ds := datastore.NewBoltDS()
	err := ds.Open()
	if err != nil {
		panic(err)
	}
	defer ds.Close()

	rrs := route.NewRepository(ds)
	rs := route.NewRouteService(rrs)

	rpc := infra.NewJSONRPCServer()
	rpc.RegisterName("RouteService", rs)

	r := infra.NewHttpRouter()

	r.GET("/", hello)
	r.POST("/rpc", rpc.ServeHTTP)

	port := "8080"
	fmt.Println("Listening on port: ", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("ListenAndServe err: ", err.Error())
	}

	fmt.Println("veloroute39: Good Buy!!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello request")
	fmt.Fprintln(w, "Hello World, I am veloroute39")
}
