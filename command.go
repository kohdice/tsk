package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newCommand() *cli.Command {
	return &cli.Command{
		Name:    "tsk",
		Version: version,
		Usage:   "Scan directories for manifest files and list executable commands",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("Hello from tsk!")
			return nil
		},
	}
}
