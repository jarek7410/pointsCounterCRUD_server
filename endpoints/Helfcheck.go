package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pointsCounterCRUD/database"
)

func halfCheck(c *gin.Context) {
	db, _ := database.Re.DB.DB()
	va := db.Ping()
	var dbStatus string
	if va != nil {
		dbStatus = "down"
	} else {
		dbStatus = "up"
	}
	data := map[string]interface{}{
		"version":  "alpha",
		"database": dbStatus,
	}
	c.JSONP(http.StatusOK, data)
}
func coffee(c *gin.Context) {
	c.String(http.StatusTeapot, "I am a teapot but programer need coffee!!!\nblik 735917147")
}
