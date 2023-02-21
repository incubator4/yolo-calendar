package pkg

import (
	"time"
)

type Calendar struct {
	ID          int       `json:"id" gorm:"primaryKey;"`
	CharacterID int       `json:"cid" gorm:"column:cid;<-"`
	Title       string    `json:"title" gorm:"<-"`
	StartTime   time.Time `json:"start_time" gorm:"column:start_time;type:time;<-"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time;type:time;<-"`
}

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UID       int    `json:"uid" gorm:"column:uid"`
	LiveID    int    `json:"live_id" gorm:"column:live_id"`
	MainColor string `json:"main_color" gorm:"column:main_color"`
}

type CharacterCalendar struct {
	Calendar
	Name      string `json:"name"`
	UID       int    `json:"uid" gorm:"column:uid"`
	LiveID    int    `json:"live_id" gorm:"column:live_id"`
	MainColor string `json:"main_color" gorm:"column:main_color"`
}

func (CharacterCalendar) TableName() string {
	return "character_calendar"
}

func (Calendar) TableName() string {
	return "calendar"
}
