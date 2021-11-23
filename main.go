package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Berglen LED",
		Usage: "Shine bright like an east london flat",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			fmt.Println("Hello world")

			return nil
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
