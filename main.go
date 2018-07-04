package main

import (
	"os"

	"github.com/po3rin/goheel/heel"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "goheel"
	app.Usage = "This app echo logs tail"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.Float64Flag{
			Name:  "num, n",
			Value: 10,
			Usage: "To specify the number of lines",
		},
		cli.BoolFlag{
			Name:  "color, c",
			Usage: "Coloring by logging level",
		},
		cli.BoolFlag{
			Name:  "watch, w",
			Usage: "watching file's change",
		},
	}
	app.Action = heel.Start
	app.Run(os.Args)
}
