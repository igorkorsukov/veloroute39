package main

import (
	"net/http"
	"os"

	"fmt"
	"handler"
	"infra"
	"route"
)

func main() {

	httproute := infra.NewHttpRouter()

	httproute.GET("/", hello)

	rs := route.NewRouteService(nil)
	rh := handler.Route{HTTP: httproute, RouteService: rs}
	rh.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Server started :", port)
	fmt.Println("error", http.ListenAndServe(":"+port, httproute).Error())
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World, I am veloroute39, привет Женя, привет Валера")
}
