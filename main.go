package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	debug   bool
	version bool

	bind string
	root string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVarP(&version, "version", "v", false, "display version information")
	flag.BoolVarP(&debug, "debug", "d", false, "enable debug logging")

	flag.StringVarP(&bind, "bind", "b", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVarP(&root, "root", "r", cwd(), "path to serve")
}

func cwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to determine current working directory: %s", err)
	}
	return cwd
}

func main() {
	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if version {
		fmt.Printf("ed version %s", FullVersion())
		os.Exit(0)
	}

	log.Fatal(NewServer(bind, root).ListenAndServe())
}
