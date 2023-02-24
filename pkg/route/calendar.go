package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/yolo-calendar/pkg"
	"github.com/incubator4/yolo-calendar/pkg/dao"
	"github.com/incubator4/yolo-calendar/pkg/types"
	"net/http"
	"time"
)

func registerCalendars(g *gin.RouterGroup) {
	g.GET("", ListCalendars)
	g.POST("")
	g.GET("/:id", ValidateCalendarID, GetCalendars)
	g.PUT("/:id", ValidateCalendarID, UpdateCalendar)
}

func ListCalendars(c *gin.Context) {
	timeRange, err := getStartAndEndOfDate(c)
	cids := c.QueryArray("cid")
	uids := c.QueryArray("uid")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	calendars, err := dao.ListCalendars(
		dao.WithUID(uids),
		dao.WithCID(cids),
		dao.WithTimeRange(timeRange),
		dao.WithOrder("id"),
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": calendars,
	})

}

func GetCalendars(c *gin.Context) {
	id := c.MustGet("id").(int)
	calendar := dao.GetCalendar(id)
	c.JSON(http.StatusOK, calendar)
}

func UpdateCalendar(c *gin.Context) {
	id := c.MustGet("id").(int)
	var calendar pkg.Calendar
	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("Calendar %d not found, err: %s", id, err),
		})
	} else {
		calendar.ID = id
		fmt.Println(calendar)
		cal := dao.UpdateCalendar(calendar)
		c.JSON(http.StatusOK, cal)
	}

}

func getStartAndEndOfDate(c *gin.Context) (types.TimeRange, error) {
	startStr := c.Query("start")
	endStr := c.Query("end")

	var start, end time.Time
	var err error

	if startStr == "" && endStr == "" {
		return types.TimeRange{}, nil
	}

	if startStr != "" && endStr != "" {
		start, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			return types.TimeRange{}, fmt.Errorf("invalid start time format, expect YYYY-MM-DD, but got %s", startStr)
		}

		end, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			return types.TimeRange{}, fmt.Errorf("invalid end time format, expect YYYY-MM-DD, but got %s", endStr)
		}

		if end.Before(start) {
			return types.TimeRange{}, fmt.Errorf("invalid time range, end time %s should not be before start time %s",
				end.Format("2006-01-02"), start.Format("2006-01-02"))
		}
	}

	if startStr != "" {
		start, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			return types.TimeRange{}, fmt.Errorf("invalid start time format, expect YYYY-MM-DD, but got %s", startStr)
		}
	}

	if endStr != "" {
		end, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			return types.TimeRange{}, fmt.Errorf("invalid end time format, expect YYYY-MM-DD, but got %s", endStr)
		}
	}
	return types.TimeRange{Start: start.Add(-8 * time.Hour), End: end.Add(-8 * time.Hour)}, nil
}
