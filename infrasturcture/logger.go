package infrasturcture

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		DisableLevelTruncation: false,
		DisableQuote:           true,
		PadLevelText:           true,
		TimestampFormat:        "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(false)
}

func ContextLogger() *log.Logger {
	// TODO 自动输出用户ID、RequestID
	return log.WithFields(log.Fields{}).Logger
}
