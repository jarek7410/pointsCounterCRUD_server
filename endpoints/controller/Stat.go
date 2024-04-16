package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pointsCounterCRUD/model"
	"pointsCounterCRUD/util"
	"strconv"
)

func CreateStat(c *gin.Context) {
	var input model.Stat
	var user_id = util.CurrentUser(c).ID

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.UserID != 0 {
		user_id = input.RoomID
	}
	booking := model.Stat{
		Value:   input.Value,
		Comment: input.Comment,
		UserID:  user_id,
		RoomID:  input.RoomID,
	}
	savedBooking, err := booking.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking done successfuly", "booking": savedBooking})
}

func GetStatsByUser(c *gin.Context) {
	var Booking []model.Stat
	var user_id = util.CurrentUser(c).ID
	err := model.GetStatByUser(user_id, &Booking)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Booking)
}

func GetStatsByRoom(c *gin.Context) {
	roomId, _ := strconv.Atoi(c.Param("id"))
	var stats []model.Stat
	//var user_id = util.CurrentUser(c).ID
	err := model.GetStatByRoom(uint(roomId), &stats)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
	c.JSON(http.StatusOK, stats)
}
