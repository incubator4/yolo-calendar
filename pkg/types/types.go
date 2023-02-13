package types

import (
	"gorm.io/gorm"
	"time"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func (tr TimeRange) DB(db *gorm.DB) *gorm.DB {
	if tr.Start.IsZero() && tr.End.IsZero() {
		return db
	} else if tr.Start.IsZero() {
		return db.Where("start_time <= ?", tr.End)
	} else if tr.End.IsZero() {
		return db.Where("start_time >= ?", tr.Start)
	} else {
		return db.Where("start_time >= ? AND start_time <= ?", tr.Start, tr.End)
	}
}
