package main

import (
	"log"

	"github.com/m99Tanishq/CLI/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
	