package main

import (
	"log"
	"os"

	"github.com/radisvaliullin/test-golang/tech/http_tcp/tcp_srv_cln_max_con/srvcln"
	"github.com/urfave/cli"
)

func main() {

	var clientNum int

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "clientnum",
				Value:       1024,
				Destination: &clientNum,
			},
		},
		Action: func(c *cli.Context) error {
			cln := srvcln.Cln{
				SrvAddr:   "0.0.0.0:7373",
				ClientNum: clientNum,
			}
			cln.Run()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
