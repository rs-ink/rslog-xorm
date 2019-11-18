package rslog_xorm

import (
	"github.com/rs-ink/rslog"
	"xorm.io/core"
)

func NewRxOrmLog() *RXOrmLog {
	return &RXOrmLog{}
}

type RXOrmLog struct {
	isDebug  bool
	logLevel core.LogLevel
	showSQL  bool
}

func (xl RXOrmLog) Debug(v ...interface{}) {
	rslog.Out(1, rslog.LevelDEBUG, v)
}

func (xl RXOrmLog) Debugf(format string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelDEBUG, format, v...)
}

func (xl RXOrmLog) Error(v ...interface{}) {
	rslog.Out(1, rslog.LevelERROR, v)
}

func (xl RXOrmLog) Errorf(format string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelERROR, format, v...)
}

func (xl RXOrmLog) Info(v ...interface{}) {
	rslog.Out(1, rslog.LevelINFO, v)
}

func (xl RXOrmLog) Infof(format string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelINFO, format, v...)
}

func (xl RXOrmLog) Warn(v ...interface{}) {
	rslog.Out(1, rslog.LevelWARN, v)
}

func (xl RXOrmLog) Warnf(format string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelWARN, format, v...)
}

func (xl RXOrmLog) Level() core.LogLevel {
	return xl.logLevel
}

func (xl *RXOrmLog) SetLevel(l core.LogLevel) {
	xl.logLevel = l
}

func (xl *RXOrmLog) ShowSQL(show ...bool) {
	if len(show) > 0 {
		xl.showSQL = show[0]
	}
}

func (xl RXOrmLog) IsShowSQL() bool {
	return xl.showSQL
}
