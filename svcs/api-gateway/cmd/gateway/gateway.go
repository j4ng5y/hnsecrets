package main

import (
	"flag"
	"log"

	"github.com/j4ng5y/hnsecrets/svcs/api-gateway/pkg/gateway/server"
)

var (
	bindAddress string
	bindPort    int
)

func parseFlags() (err error) {
	flag.StringVar(&bindAddress, "address", "0.0.0.0", "The IP Address to bind to.")
	flag.IntVar(&bindPort, "port", 8080, "The TCP port to bind to.")
	flag.Parse()
	return nil
}

func main() {
	if err := parseFlags(); err != nil {
		log.Fatal(err)
	}

	S, err := server.New(server.WithDefaults())
	if err != nil {
		log.Fatal(err)
	}

	S.Run()
}
