package main

import (
	"os"

	"github.com/rotationalio/subscribr"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "subscribr"
	app.Usage = "runs the subscriber server or interacts with it"
	app.Version = subscribr.Version()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "S, no-secure",
			Usage: "don't connect with TLS to the server",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "serve",
			Usage:    "run the subscribr server",
			Category: "server",
			Action:   serve,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "a, addr",
					Usage:  "address to bind the server on",
					EnvVar: "SUBSCRIBR_BIND_ADDR",
					Value:  ":2383",
				},
			},
		},
		{
			Name:     "subscribe",
			Usage:    "subscribe to subscribr",
			Category: "client",
			Action:   subscribe,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "e, email",
					Usage: "specify the email to subscribe with",
				},
			},
		},
	}

	app.Run(os.Args)
}

func serve(c *cli.Context) (err error) {
	addr := c.String("addr")
	if addr == "" {
		return cli.NewExitError("must specify an address to bind on", 1)
	}

	return nil
}

func subscribe(c *cli.Context) (err error) {
	return cli.NewExitError("not implemented yet", 43)
}
