package main

import (
	"github.com/mschlech/SimpleJsonLogger/pkg/app"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	application := cli.NewApp()

	application.Name = "simple-json-logger"
	application.Usage = "Start json logger"

	log.WithFields(log.Fields{"package": "main"}).Info("Starting jsonLogger")

	application.Action = func(c *cli.Context) error {
		return app.StartJsonLogger()
	}
	application.Run(os.Args)
}
