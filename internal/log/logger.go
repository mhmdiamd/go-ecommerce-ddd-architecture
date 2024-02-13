package log

import "github.com/NooBeeID/go-logging/logger"

var Log logger.Logger

func Init() {
  Log = logger.NewLog()
  Log.SetReportCaller(true)
}
