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

	//if !infra.IsGCloud() {
	//	cr := filepath.Join("c:", "gcloud", "veloroute39-78759ec9aa23.json")
	//	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cr)
	//}
	//
	//ds := NewDS()
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
	r.GET("/_ah/health", health)
	r.POST("/rpc", rpc.ServeHTTP)

	port := "8080"
	fmt.Println("Listening on port: ", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("ListenAndServe err: ", err.Error())
	}

	fmt.Println("veloroute39: Good Buy!!")
}

func NewDS() datastore.DataStore {
	if !infra.IsGCloud() {
		return datastore.NewBoltDS()
	}
	return nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello request")
	fmt.Fprintln(w, "Hello World, I am veloroute39")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("health request")
	fmt.Fprint(w, "ok")
}

//gcloud.cmd app deploy --project veloroute39
