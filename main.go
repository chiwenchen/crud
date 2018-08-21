package main

import (
	"errors"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// User is the reflection of User table
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	RegionId int    `json:"region_id"`
}

var errEmpty = errors.New("empty string")

func main() {
	svc := userService{}

	fetchUserHandler := httptransport.NewServer(
		FetchUserEndpoint(svc),
		DecodeFetchUserRequest,
		EncodeFetchUserResponse,
	)

	createUserHandler := httptransport.NewServer(
		CreateUserEndpoint(svc),
		DecodeCreateUserRequest,
		EncodeFetchUserResponse,
	)

	updateUserHandler := httptransport.NewServer(
		UpdateUserEndpoint(svc),
		DecodeUpdateUserRequest,
		EncodeFetchUserResponse,
	)

	deleteUserHandler := httptransport.NewServer(
		DeleteUserEndpoint(svc),
		DecodeFetchUserRequest,
		EncodeFetchUserResponse,
	)

	http.Handle("/find_user", fetchUserHandler)
	http.Handle("/create_user", createUserHandler)
	http.Handle("/update_user", updateUserHandler)
	http.Handle("/delete_user", deleteUserHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
