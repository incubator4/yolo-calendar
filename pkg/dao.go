package pkg

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

func ListCalendars(params ListCalendarParams) []CharacterCalendar {
	var calendars []CharacterCalendar
	db := params.TimeRange.DB(DB)
	if len(params.CIDArray) > 0 {
		db.Where("cid IN (?)", params.CIDArray).Find(&calendars)
	} else if len(params.UIDArray) > 0 {
		db.Where("uid IN (?)", params.UIDArray).Find(&calendars)
	} else {
		db.Find(&calendars)
	}

	return calendars
}

func GetCharacter(_c Character) *Character {
	var c = new(Character)
	DB.Table("character").Where(&_c).First(c)
	return c
}

func ListCharacter() []Character {
	var characters []Character
	DB.Table("character").Order("id").Find(&characters)
	return characters
}
