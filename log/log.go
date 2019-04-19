package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}
