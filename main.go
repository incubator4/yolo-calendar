package main

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/yolo-calendar/pkg"
	"github.com/incubator4/yolo-calendar/pkg/config"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println(config.GlobalConfig)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	r := gin.Default()
	r.GET("/api/uid/:uid", func(c *gin.Context) {
		stringUID := c.Param("uid")
		uid, err := strconv.Atoi(stringUID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
			return
		}
		character := pkg.GetCharacter(pkg.Character{UID: uid})
		c.JSON(http.StatusOK, gin.H{
			"data": character,
		})
	})
	r.GET("/api/ics/:mixId", func(c *gin.Context) {
		nameOrUID := c.Param("mixId")
		var params pkg.ListCalendarParams
		if nameOrUID == "team" {
			params.All = true
		} else {
			var character = new(pkg.Character)
			intID, err := strconv.Atoi(nameOrUID)
			if err != nil {
				//c.JSON(http.StatusOK, gin.H{
				//	"error": err,
				//})
				character = pkg.GetCharacter(pkg.Character{Name: nameOrUID})
			} else {
				character = pkg.GetCharacter(pkg.Character{UID: intID})
			}

			params.ID = character.ID

		}
		calendars := pkg.ListCalendars(params)
		cal := ics.NewCalendar()
		cal.SetMethod(ics.MethodRequest)

		for _, calendar := range calendars {
			realTime := calendar.DateTime.Add(-8 * time.Hour).In(loc)
			e := ics.NewEvent(fmt.Sprintf("%d - %s", calendar.CharacterID, realTime))
			e.SetCreatedTime(realTime)
			e.SetStartAt(realTime)
			e.SetSummary(calendar.Title)
			cal.AddVEvent(e)
		}
		c.String(http.StatusOK, cal.Serialize())
	})
	_ = r.Run()
	//cal := ics.NewCalendar()
	//cal.SetMethod(ics.MethodRequest)
	//event := cal.AddEvent(fmt.Sprintf("id@domain", 1))
	//event.SetCreatedTime(time.Now())
	//event.SetDtStampTime(time.Now())
	//event.SetModifiedAt(time.Now())
	//event.SetStartAt(time.Now())
	//event.SetEndAt(time.Now())
	//event.SetSummary("Summary")
	//event.SetLocation("Address")
	//event.SetDescription("Description")
	//event.SetURL("https://URL/")
	//event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
	//event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
	//event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
	//fmt.Println(cal.Serialize())
	//var calendars []pkg.Calendar
	//pkg.DB.Table("calendar").Find(&calendars, 3)
	//fmt.Println(calendars)

}
