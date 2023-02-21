package dao

import "github.com/incubator4/yolo-calendar/pkg"

func ListCalendars(params ListCalendarParams) []pkg.CharacterCalendar {
	var calendars []pkg.CharacterCalendar
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

func GetCalendar(id int) *pkg.CharacterCalendar {
	var c = new(pkg.CharacterCalendar)
	DB.First(&c, id)
	return c
}

func UpdateCalendar(cal pkg.Calendar) *pkg.CharacterCalendar {

	DB.Save(&cal)
	var c = new(pkg.CharacterCalendar)
	DB.Where("id = ?", cal.ID).First(&c)
	return c
}
