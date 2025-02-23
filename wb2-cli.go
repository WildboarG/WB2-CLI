package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/wb2-cli/command"
)

func main() {
	app := &cli.App{
		Name:    "wb2-cli",
		Usage:   "wb2-cli [commands]",
		Version: "v0.0.1",
		Commands: []*cli.Command{
			command.Init(),
			// command.Version(),
			// command.Config(),
			command.Create(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
