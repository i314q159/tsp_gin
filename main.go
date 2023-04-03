package main

import (
	"log"
	"net/http"
	"tsp_gin/router"

	"golang.org/x/sync/errgroup"
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

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
