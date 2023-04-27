package main

import (
	"fmt"
	"log"
	"net/http"
	"tsp_gin/conf"
	"tsp_gin/router"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	serverA := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.SERVER_IP, conf.SERVER_PORT),
		Handler: router.TspRouter(),
	}

	g.Go(func() error {
		return serverA.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
