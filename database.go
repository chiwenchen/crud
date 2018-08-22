package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User is the reflection of User table
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	RegionID int    `json:"region_id"`
}

// InitDB to interact with Mysql
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root@unix(/tmp/mysql.sock)/snapask_development")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}
