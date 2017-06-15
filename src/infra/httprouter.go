package infra

import (
	"net/http"

	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type Middleware func(http.Handler) http.Handler
type RouteParam struct {
	Key   string
	Value string
}
type RouteParams []RouteParam

func (ps RouteParams) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

type HttpRouter interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)

	Use(m ...Middleware)

	PUT(path string, hf http.HandlerFunc)
	HEAD(path string, hf http.HandlerFunc)
	GET(path string, hf http.HandlerFunc)
	POST(path string, hf http.HandlerFunc)
	DELETE(path string, hf http.HandlerFunc)
}

type router struct {
	router *httprouter.Router
	chain  alice.Chain
}

func NewHttpRouter() HttpRouter {
	return &router{router: httprouter.New()}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *router) Use(mws ...Middleware) {
	acs := make([]alice.Constructor, 0)
	for _, mw := range mws {
		acs = append(acs, alice.Constructor(mw))
	}
	r.chain = r.chain.Append(acs...)
}

func (r *router) PUT(path string, hf http.HandlerFunc) {
	r.router.PUT(path, r.wrap(hf))
}

func (r *router) HEAD(path string, hf http.HandlerFunc) {
	r.router.HEAD(path, r.wrap(hf))
}

func (r *router) GET(path string, hf http.HandlerFunc) {
	r.router.GET(path, r.wrap(hf))
}

func (r *router) POST(path string, hf http.HandlerFunc) {
	r.router.POST(path, r.wrap(hf))
}

func (r *router) DELETE(path string, hf http.HandlerFunc) {
	r.router.DELETE(path, r.wrap(hf))
}

func (r *router) wrap(h http.HandlerFunc) httprouter.Handle {
	return r.stripParams(r.chain.ThenFunc(h))
}

func (r *router) stripParams(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, hps httprouter.Params) {

		ps := RouteParams{}
		for _, p := range hps {
			ps = append(ps, RouteParam{Key: p.Key, Value: p.Value})
		}

		ctx := context.WithValue(r.Context(), "routeparams", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}
