package pkg

import (
	"github.com/incubator4/yolo-calendar/pkg/config"
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
	ID  int
	All bool
}

func ListCalendars(params ListCalendarParams) []Calendar {
	var calendars []Calendar
	if params.All {
		DB.Table("calendar").Find(&calendars)
	} else {
		DB.Table("calendar").Find(&calendars, params.ID)
	}
	return calendars
}

func GetCharacter(_c Character) *Character {
	var c = new(Character)
	DB.Table("character").Where(&_c).First(c)
	return c
}
