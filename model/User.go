package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"pointsCounterCRUD/database"
	"strings"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	RoleID   uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// Save user details
func (user *User) Save() (*User, error) {
	err := database.Re.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// Validate user password
func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Get all users
func GetUsers(User *[]User) (err error) {
	err = database.Re.DB.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := database.Re.DB.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func GetUserById(id uint) (User, error) {
	var user User
	err := database.Re.DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func GetUser(User *User, id int) (err error) {
	err = database.Re.DB.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(User *User) (err error) {
	err = database.Re.DB.Omit("password").Updates(User).Error
	if err != nil {
		return err
	}
	return nil
}
