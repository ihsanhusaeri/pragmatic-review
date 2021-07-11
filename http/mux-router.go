package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatch = mux.NewRouter()
)

func (r *muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatch.HandleFunc(uri, f).Methods("POST")
}

func (r *muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatch.HandleFunc(uri, f).Methods("GET")
}
func (r *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatch)
}
