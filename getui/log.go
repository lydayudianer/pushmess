package getui

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
}
