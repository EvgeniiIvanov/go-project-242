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
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			options := goproject242.Options{
				Recursive: cmd.Bool("recursive"),
				All:       cmd.Bool("all"),
			}
			size, err := goproject242.GetSize(cmd.Args().Get(0), options)
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
