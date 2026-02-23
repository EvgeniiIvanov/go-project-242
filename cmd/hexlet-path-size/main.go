package main

import (
	"context"
	"fmt"
	"log"
	"os"

	goproject242 "code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "Count size of files and directories",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			size, err := goproject242.GetSize(cmd.Args().Get(0))
			if err != nil {
				return err
			}
			if cmd.Bool("human") {
				fmt.Printf("%v\\%s\n", goproject242.FormatSize(size, true), cmd.Args().Get(0))
				return nil
			}
			fmt.Printf("%vB\\%s\n", size, cmd.Args().Get(0))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
