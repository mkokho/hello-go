package main

import (
	"flag"
	"log"
	"os"
)

type config struct {
  dbFile string
}

func mustParseConfig() config {
	filename := flag.String("db", "", "db file of the storage")

	flag.Parse()
	if *filename == "" {
		log.Print("db config is required")
		flag.Usage()
		os.Exit(1)
	}

	return config{
		dbFile: *filename,
	}

}
