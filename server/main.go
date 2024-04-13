package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	connect "connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	countv1 "github.com/ahmad-ibra/count-server/gen/count/v1"
	"github.com/ahmad-ibra/count-server/gen/count/v1/countv1connect"
)

const address = "0.0.0.0"

var count int32 = 0

func main() {
	mux := http.NewServeMux()
	path, handler := countv1connect.NewCountServiceHandler(&countServiceServer{})
	mux.Handle(path, handler)
	port := os.Getenv("PORT")
	fmt.Printf("Listening on %v:%v...\n", address, port)
	http.ListenAndServe(
		address+":"+port,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// countServiceServer implements the CountService API.
type countServiceServer struct {
	countv1connect.UnimplementedCountServiceHandler
}

// Count counts the number of times it has been called.
func (s *countServiceServer) Count(ctx context.Context, req *connect.Request[countv1.CountRequest]) (*connect.Response[countv1.CountResponse], error) {
	count++
	name := req.Msg.Name

	log.Printf("Got a request from %s to count: %d", name, count)
	return connect.NewResponse(&countv1.CountResponse{Count: count}), nil
}
