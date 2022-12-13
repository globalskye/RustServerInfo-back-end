package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogrus() {
	file, err := os.OpenFile("logs/log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		logrus.Fatalf("failed to open log file:%s", err.Error())
	}
	logrus.SetOutput(file)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(new(logrus.JSONFormatter))
}
