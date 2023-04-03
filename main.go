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
	server_A := &http.Server{
		Addr:    ":8080",
		Handler: router.Router_A(),
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	server_B := &http.Server{
		Addr:    ":8081",
		Handler: router.Router_B(),
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server_A.ListenAndServe()
	})

	g.Go(func() error {
		return server_B.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
