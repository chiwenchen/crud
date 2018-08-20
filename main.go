package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// InitDB to interact with Mysql
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root@unix(/tmp/mysql.sock)/snapask_development")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}

// User is the reflection of User table
type User struct {
	ID       string
	Username string
}

var errEmpty = errors.New("empty string")

// UserService is for user table interaction
type UserService interface {
	FetchUser(id string) (User, error)
	// CreateUser *User
	// UpdateUser(id string) *User
	// DeleteUser(id string) *User
}

type userService struct{}

func (userService) FetchUser(id string) (User, error) {
	if id == "" {
		return User{}, errEmpty
	}

	var user User

	db, _ := InitDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, nil
	} else {
		return user, nil
	}
}

type fetchUserRequest struct {
	ID string `json:"id"`
}
type fetchUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Err      string
}

func fetchUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(fetchUserRequest)
		v, err := svc.FetchUser(req.ID)
		if err != nil {
			return fetchUserResponse{v.ID, v.Username, "encode to response failed"}, nil
		}
		return fetchUserResponse{v.ID, v.Username, ""}, nil

	}
}

func decodeFetchUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req fetchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeFetchUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := userService{}

	fetchUserHandler := httptransport.NewServer(
		fetchUserEndpoint(svc),
		decodeFetchUserRequest,
		encodeFetchUserResponse,
	)

	http.Handle("/fetch_user", fetchUserHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
