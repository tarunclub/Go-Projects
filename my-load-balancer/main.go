package main

import (
	"log"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newServer(addr string) *Server {
	serverAddr, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverAddr),
	}
}
