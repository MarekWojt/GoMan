package http

import (
	"github.com/valyala/fasthttp"
)

// Run runs the http server
func Run() error {
	return fasthttp.ListenAndServeTLS(
		"localhost:54321",
		"certs/localhost.crt",
		"certs/localhost.key",
		handleRequest,
	)
}
