package pkg

import (
	"time"
)

type Calendar struct {
	CharacterID int       `json:"cid" gorm:"column:cid;<-:false"`
	Title       string    `json:"title" gorm:"<-:false"`
	DateTime    time.Time `json:"dateTime" gorm:"column:datetime;type:time;<-:false"`
}

type Character struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UID    int    `json:"uid" gorm:"column:uid"`
	LiveID int    `json:"live_id" gorm:"column:live_id"`
}
