package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"pkg.nimblebun.works/tsc-language-server/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "tsc-ls"
	app.Usage = "language Server for the TSC scripting language"
	app.Version = "0.1.6"

	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "start the language server in stdio mode",
			Action: commands.StartCommand,
		},
		{
			Name:   "tcp",
			Usage:  "start the language server in TCP mode",
			Action: commands.TCPCommand,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "port",
					Value: 17881,
					Usage: "port on which the TCP server should start",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
