package simpleJsonLogger

import (
	"github.com/mschlech/SimpleJsonLogger/private/app/simple-json-logger/lib"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "simple-json-logger"
	app.Usage = "Start json logger"

	log.WithFields(log.Fields{"package": "main"}).Info("Starting jsonLogger")

	app.Action = func() error {
		return lib.StartJsonLogger()
	}
	app.Run(os.Args)
}
