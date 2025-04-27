package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var version = "0.0.1"

func main() {
	cmd := &cli.Command{
		Name:    "tsk",
		Version: version,
		Usage:   "Scan directories for manifest files and list executable commands",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("Hello from tsk!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
