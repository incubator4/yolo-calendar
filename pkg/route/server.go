package route

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/yolo-calendar/pkg/dao"
	"net/http"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		cal := api.Group("/cal")
		registerCalendars(cal)
	}

	{
		vtuber := api.Group("/vtubers")
		characters := api.Group("/characters")
		registerVtubers(vtuber)
		registerVtubers(characters)
	}

	r.GET("/api/ics/:uid", func(c *gin.Context) {
		UID := c.Param("uid")
		var params = dao.ListCalendarParams{
			UIDArray: []string{UID},
		}
		calendars := dao.ListCalendars(params)

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

	return r
}
