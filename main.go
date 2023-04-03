package main

import (
	"golang.org/x/sync/errgroup"
	"net/http"
	"tsp_gin/router"
)

var (
	g errgroup.Group
)

func main() {

	serverA := &http.Server{
		Addr:    ":8080",
		Handler: router.TspRouter(),
	}

	g.Go(func() error {
		return serverA.ListenAndServe()
	})
}
