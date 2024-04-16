package model

import "pointsCounterCRUD/database"

type Stat struct {
	ID      uint   `gorm:"primary_key"`
	Value   int    `gorm:"not null" json:"value"`
	Comment string `json:"comment"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	RoomID  uint   `gorm:"not null" json:"room_id"`
	User    User   `json:"-"`
	Room    Room   `json:"-"`
}

func GetStatByUser(userId uint, stat *[]Stat) error {
	//var stat []Stat
	err := database.Re.DB.Where("user_id=?", userId).Find(&stat).Error
	if err != nil {
		return err
	}
	return nil
}
func GetStatByRoom(roomId uint, stat *[]Stat) error {
	//var stat []Stat
	err := database.Re.DB.Where("room_id=?", roomId).Find(&stat).Error
	if err != nil {
		return err
	}
	return nil
}

func (stat *Stat) Save() (*Stat, error) {
	err := database.Re.DB.Create(&stat).Error
	if err != nil {
		return nil, err
	}
	return stat, nil
}
func (stat *Stat) Erace() (*Stat, error) {
	err := database.Re.DB.Delete(&stat).Error
	if err != nil {
		return stat, err
	}
	return stat, nil
}
