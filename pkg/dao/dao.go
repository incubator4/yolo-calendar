package dao

import (
	"github.com/incubator4/yolo-calendar/pkg/config"
	"github.com/incubator4/yolo-calendar/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := config.GlobalConfig.DbUrl
	DB, err = gorm.Open(postgres.New(postgres.Config{

		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		panic(err)
	}

}

type ListCalendarParams struct {
	CIDArray  []string
	UIDArray  []string
	TimeRange types.TimeRange
}

type Option func(*gorm.DB) *gorm.DB

func WithTimeRange(tr types.TimeRange) Option {
	return func(db *gorm.DB) *gorm.DB {
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
}

func WithUID(UIDArray []string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if UIDArray != nil && len(UIDArray) > 0 {
			return db.Where("uid IN (?)", UIDArray)
		} else {
			return db
		}
	}
}

func WithCID(CIDArray []string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if CIDArray != nil && len(CIDArray) > 0 {
			return db.Where("cid IN (?)", CIDArray)
		} else {
			return db
		}
	}
}

func WithOrder(order string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithNotDelete() Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_delete = ?", false)
	}
}

func WithActive() Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true)
	}
}

func WithOwner(id int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("own_id = ?", id)
	}
}
