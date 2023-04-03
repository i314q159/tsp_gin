package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"tsp_gin/router"
)

var (
	g errgroup.Group
)

func main() {
	serverA := &http.Server{
		Addr:    ":8080",
		Handler: router.RouterA(),
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	serverB := &http.Server{
		Addr:    ":8081",
		Handler: router.RouterB(),
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return serverA.ListenAndServe()
	})

	g.Go(func() error {
		return serverB.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
