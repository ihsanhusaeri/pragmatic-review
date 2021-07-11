package http

import "net/http"

type Router interface {
	POST(uri string, f func(resp http.ResponseWriter, req *http.Request))
	GET(uri string, f func(resp http.ResponseWriter, req *http.Request))
	SERVE(port string)
}
