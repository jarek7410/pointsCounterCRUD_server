package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func halfCheck(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
func coffee(c *gin.Context) {
	c.String(http.StatusTeapot, "I am a teapot but programer need coffee!!!")
}
