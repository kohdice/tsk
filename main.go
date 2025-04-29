package main

import (
	"context"
	"log"
	"os"
)

var version = "0.0.1"

func main() {
	cmd := newCommand()
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
