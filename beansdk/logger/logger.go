package logger

import (
	"github.com/anjude/bcore/pkg/logger"
)

var log = logger.Default()

// Define logrus alias
var (
	Debugf = log.Debugf
	Infof  = log.Infof
	Warnf  = log.Warnf
	Errorf = log.Errorf
	Fatalf = log.Fatalf
	Panicf = log.Panicf
	Printf = log.Printf
	Info   = log.Info
	Debug  = log.Debug
	Error  = log.Error
)

func GetLogger(traceId string) *logger.BeanLogger {
	newLog := log
	newLog.SetFormatter(&logger.BeanFormatter{TraceId: traceId})
	return newLog
}
