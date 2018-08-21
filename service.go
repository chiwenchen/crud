package main

import "errors"

// UserService is for user table interaction
type UserService interface {
	FetchUser(id string) (*User, error)
	CreateUser(username string, rId int) (*User, error)
	UpdateUser(id string, username string) (*User, error)
	// DeleteUser(id string) *User
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
	} else {
		return &user, nil
	}
}
func (userService) CreateUser(username string, rId int) (*User, error) {

	user := User{Username: username, RegionId: rId}

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
