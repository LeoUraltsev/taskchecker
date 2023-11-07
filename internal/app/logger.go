package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func SetLogrus(level string) {
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(parseLevel)
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.DateTime,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			s := strings.Split(frame.Function, ".")
			function = fmt.Sprintf("Function = %s()", s[len(s)-1])
			dir, f := path.Split(frame.File)
			file = fmt.Sprintf(" Path = %s%s", dir, f)
			return
		},
	})

	logrus.SetOutput(os.Stdout)
}
