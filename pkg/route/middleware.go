package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Middleware func(handler gin.HandlerFunc) gin.HandlerFunc

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	//c.Request.Cookies()
	if token == "example-token" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
}

func ValidateCalendarID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("Invalid calendar ID %s", err),
		})

	} else {
		c.Set("id", id)
		c.Next()
	}
}
