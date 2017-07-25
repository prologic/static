package main

import (
	"log"
	"os"

	"github.com/namsral/flag"
)

func cwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to determine current working directory: %s", err)
	}
	return cwd
}

func main() {
	var (
		bind string
		root string
	)

	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVar(&root, "root", cwd(), "path to serve")
	flag.Parse()

	NewServer(bind, root).ListenAndServe()
}
