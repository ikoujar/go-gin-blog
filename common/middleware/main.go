package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoggerMiddleware(c *gin.Context) {
	// agent := c.Request.Header["User-Agent"]
	// fmt.Println(agent)
	c.Next()
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//do something here
		c.Next()
		//do something base on error in c.
		for _, err := range c.Errors {
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
			}
		}

		//else continue regular processing.
	}
}
