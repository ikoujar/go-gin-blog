package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ThrowErrorIfExist(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
		return true // signal that there was an error and the caller should return
	}
	return false // no error, can continue
}
