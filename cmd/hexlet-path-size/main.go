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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			size, err := goproject242.GetSize(cmd.Args().Get(0))
			if err != nil {
				return err
			}
			fmt.Printf("%vB\\%s\n", size, cmd.Args().Get(0))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
