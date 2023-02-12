package pkg

import (
	"time"
)

type Calendar struct {
	CharacterID int       `json:"cid" gorm:"column:cid;<-:false"`
	Title       string    `json:"title" gorm:"<-:false"`
	StartTime   time.Time `json:"start_time" gorm:"column:start_time;type:time;<-:false"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time;type:time;<-:false"`
}

type Character struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UID    int    `json:"uid" gorm:"column:uid"`
	LiveID int    `json:"live_id" gorm:"column:live_id"`
}

type CharacterCalendar struct {
	Calendar
	Name   string `json:"name"`
	UID    int    `json:"uid" gorm:"column:uid"`
	LiveID int    `json:"live_id" gorm:"column:live_id"`
}

func (CharacterCalendar) TableName() string {
	return "character_calendar"
}
