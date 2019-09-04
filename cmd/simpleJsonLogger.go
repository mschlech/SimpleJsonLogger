package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"github.com/mschlech/SimpleJsonLogger/private/app/simpleJsonLogger/
)

func main() {
	app := cli.NewApp()

	app.Name = "simpleJsonLogger"
	app.Usage = "Start json logger"

	log.WithFields(log.Fields{"package": "main"}).Info("Starting jsonLogger")
	return nil
	os.Args()
}
