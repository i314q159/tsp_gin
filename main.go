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

	serverB := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.WEB_IP, conf.WEB_PORT),
		Handler: router.WebRouter(),
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
