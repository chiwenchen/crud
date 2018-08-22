package main

import "errors"

// UserService is for user table interaction
type UserService interface {
	FetchUser(id string) (*User, error)
	CreateUser(username string, regionID int) (*User, error)
	UpdateUser(id string, username string) (*User, error)
	DeleteUser(id string) (*User, error)
}

type userService struct{}

func (userService) FetchUser(id string) (*User, error) {
	if id == "" {
		return &User{}, errEmpty
	}

	var user User

	db, _ := InitDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}
func (userService) CreateUser(username string, regionID int) (*User, error) {

	user := User{Username: username, RegionID: regionID}

	db, _ := InitDB()

	db.Create(&user)

	return &user, errors.New("")
}

func (userService) UpdateUser(id string, username string) (*User, error) {
	db, _ := InitDB()

	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return &User{}, err
	}

	user.Username = username
	db.Save(&user)

	return &user, nil
}

func (userService) DeleteUser(id string) (*User, error) {
	db, _ := InitDB()

	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return &User{}, err
	}

	db.Delete(&user)

	return &user, nil
}
