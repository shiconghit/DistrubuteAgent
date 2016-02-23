package main

import (
	"./g"
	"./http"
	"flag"
	"fmt"
	"log"
	"os"
	"./mysql"
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

	mysql.InitDb()

	go http.Start()

	select {}
}
