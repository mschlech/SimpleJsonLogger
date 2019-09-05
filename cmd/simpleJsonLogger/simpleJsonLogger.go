package simpleJsonLogger

import (
	"github.com/mschlech/SimpleJsonLogger/private/app/simpleJsonLogger"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "simpleJsonLogger"
	app.Usage = "Start json logger"

	log.WithFields(log.Fields{"package": "main"}).Info("Starting jsonLogger")

	app.Action = func() error {
		return simpleJsonLogger.StartJsonLogger()
	}
	app.Run(os.Args)
}
