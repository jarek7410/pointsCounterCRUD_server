package query

import (
	"pointsCounterCRUD/database"
)

// Create a role
func CreateRole(Role *Role) (err error) {
	err = database.DB.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func GetRoles(Role *[]Role) (err error) {
	err = database.DB.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func GetRole(Role *Role, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func UpdateRole(Role *Role) (err error) {
	database.DB.Save(Role)
	return nil
}
