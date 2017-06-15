package infra

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// adapt HTTP connection to ReadWriteCloser
type httpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *httpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *httpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *httpConn) Close() error                      { return nil }

type JSONRPCServer struct {
	*rpc.Server
}

func NewJSONRPCServer() *JSONRPCServer {
	return &JSONRPCServer{rpc.NewServer()}
}

func (s *JSONRPCServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("rpc: request")
	codec := jsonrpc.NewServerCodec(&httpConn{in: r.Body, out: w})
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	err := s.Server.ServeRequest(codec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, details have been logged.", 500)
		return
	}
}
