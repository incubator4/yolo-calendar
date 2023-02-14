package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/yolo-calendar/pkg"
	"github.com/incubator4/yolo-calendar/pkg/types"
	"net/http"
	"time"
)

func ListCalendars(c *gin.Context) {
	timeRange, err := getStartAndEndOfDate(c)
	cids := c.QueryArray("cid")
	uids := c.QueryArray("uid")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		calendars := pkg.ListCalendars(pkg.ListCalendarParams{UIDArray: uids, CIDArray: cids, TimeRange: timeRange})
		c.JSON(http.StatusOK, gin.H{
			"data": calendars,
		})
	}

}

func getStartAndEndOfDate(c *gin.Context) (types.TimeRange, error) {
	startStr := c.Query("start")
	endStr := c.Query("end")

	var start, end time.Time
	var err error

	if startStr == "" && endStr == "" {
		return getCurrentWeekTimeRange(), nil
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

func getCurrentWeekTimeRange() types.TimeRange {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	start := now.AddDate(0, 0, 1-weekday)
	end := start.AddDate(0, 0, 7)
	return types.TimeRange{start, end}
}
