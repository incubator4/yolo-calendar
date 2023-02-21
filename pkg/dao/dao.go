package dao

import (
	"github.com/incubator4/yolo-calendar/pkg/config"
	"github.com/incubator4/yolo-calendar/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

var Loc *time.Location

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

	Loc, _ = time.LoadLocation("Asia/Shanghai")
}

type ListCalendarParams struct {
	CIDArray  []string
	UIDArray  []string
	TimeRange types.TimeRange
}
