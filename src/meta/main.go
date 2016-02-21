package main

import (
	"flag"
	"fmt"
	"./g"
	"./http"
	"os"
    "log"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if err := g.ParseConfig(*cfg); err != nil {
        log.Println(err)
	}

	go http.Start()

	select {}
}
