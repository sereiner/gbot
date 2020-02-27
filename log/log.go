package log

import (
	"github.com/sereiner/gbot/gbot"
	logger "github.com/sereiner/library/log"
	"github.com/wule61/log"
	"os"
)

type Log struct {
	gbot gbot.Gbot
	logx *logger.Logger
	log  logger.ILogging
}

func NewLog(gbot gbot.Gbot) interface{} {
	logging := log.New(os.Stdout, "", log.Llongcolor)
	logging.SetOutputLevel(log.Ldebug)
	return &Log{
		gbot: gbot,
		logx: logger.GetSession("message", logger.CreateSession()),
		log:  logging,
	}
}

func (l *Log) Logx() *logger.Logger {
	return l.logx
}

func (l *Log) Log() logger.ILogging {
	return l.log
}
