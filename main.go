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

	g.Go(func() error {
		return serverA.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
