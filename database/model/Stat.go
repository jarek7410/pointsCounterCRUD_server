package model

type Stat struct {
	ID    uint `gorm:"primary_key"`
	Value int  `json:"value"`
	User  User `json:"user"`
}
