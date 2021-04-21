package main

import (
	"fmt"
	"io"
	"log"
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
	log.Printf("%s %s %s", r.Method, r.URL.String(), r.Proto)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%s %s %s\r\n", r.Method, r.RequestURI, r.Proto)
	fmt.Fprintf(w, "Host: %s\r\n", r.Host)
	r.Header.Write(w)
	fmt.Fprintf(w, "\r\n")
	io.Copy(w, r.Body)
}
