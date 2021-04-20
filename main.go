package main

import (
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var h http.Handler
	h = http.HandlerFunc(echo)
	h = h2c.NewHandler(h, &http2.Server{})

	http.ListenAndServe(":"+port, h)
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r.Write(w)
}
