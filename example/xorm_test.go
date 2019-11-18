package example

import (
	"github.com/rs-ink/rslog-xorm"
	"xorm.io/xorm"
)

var db *xorm.Engine

func init() {
	if db != nil {
		db.SetLogger(rslog_xorm.NewRxOrmLog())
	}
}
