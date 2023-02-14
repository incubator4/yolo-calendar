package route

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/yolo-calendar/pkg"
	"net/http"
	"strconv"
)

func NewServer() *gin.Engine {
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
	r.GET("/api/ics/:uid", func(c *gin.Context) {
		UID := c.Param("uid")
		var params = pkg.ListCalendarParams{
			UIDArray: []string{UID},
		}
		calendars := pkg.ListCalendars(params)

		cal := ics.NewCalendar()
		cal.SetMethod(ics.MethodRequest)
		cal.SetTzid("Asia/Shanghai")

		for _, calendar := range calendars {
			e := ics.NewEvent(fmt.Sprintf("%d - %s", calendar.CharacterID, calendar.StartTime))
			e.SetCreatedTime(calendar.StartTime)
			e.SetStartAt(calendar.StartTime)
			e.SetEndAt(calendar.EndTime)
			e.SetSummary(fmt.Sprintf("%s %s", calendar.Name, calendar.Title))
			cal.AddVEvent(e)
		}
		c.String(http.StatusOK, cal.Serialize())
	})
	r.GET("/api/cal", ListCalendars)

	return r
}
