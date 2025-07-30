package main

import (
	"log"

	"github.com/m99Tanishq/glm-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
