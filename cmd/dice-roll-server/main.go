package main

import (
	"fmt"
	"github.com/hamologist/dice-roll/pkg/app"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

func main() {
	router := app.NewRouter()

	cliApp := &cli.App{
		Flags: []cli.Flag {
			&cli.IntFlag{
				Name:        "port",
				Usage:       "port to run the server on",
				Value:       3000,
			},
			&cli.StringFlag{
				Name:        "host",
				Usage:       "host to run the server on",
				Value:       "",
			},
		},
		Name: "dice-roll-server",
		Usage: "starts a server for processing dice roll requests",
		Action: func(c *cli.Context) error {
			return http.ListenAndServe(
				fmt.Sprintf("%s:%d", c.String("host"), c.Int("port")),
				router,
			)
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}