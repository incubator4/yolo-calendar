package pkg

import (
	"github.com/incubator4/yolo-calendar/pkg/config"
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
	CID int
	All bool
}

func ListCalendars(params ListCalendarParams) []CharacterCalendar {
	var calendars []CharacterCalendar
	if params.All {
		DB.Find(&calendars)
	} else {
		DB.Find(&calendars, params.CID)
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
	DB.Table("character").Find(&characters)
	return characters
}
