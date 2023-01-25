package main

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/yolo-calendar/pkg"
	"github.com/incubator4/yolo-calendar/pkg/config"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println(config.GlobalConfig)

	r := gin.Default()
	r.GET("/api/characters", func(c *gin.Context) {

		characters := pkg.ListCharacter()
		c.JSON(http.StatusOK, gin.H{
			"data": characters,
		})
	})
	r.GET("/api/characters/:uid", func(c *gin.Context) {
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
		cal.SetTzid("Asia/Shanghai")

		for _, calendar := range calendars {
			e := ics.NewEvent(fmt.Sprintf("%d - %s", calendar.CharacterID, calendar.DateTime))
			e.SetCreatedTime(calendar.DateTime)
			e.SetStartAt(calendar.DateTime)
			e.SetSummary(calendar.Title)
			cal.AddVEvent(e)
		}
		c.String(http.StatusOK, cal.Serialize())
	})
	r.GET("/api/cal", func(c *gin.Context) {

		calendars := pkg.ListCalendars(pkg.ListCalendarParams{All: true})
		c.JSON(http.StatusOK, gin.H{
			"data": calendars,
		})
	})
	r.GET("/api/cal/:mixId", func(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{
			"data": calendars,
		})
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
