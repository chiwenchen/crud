package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type fetchUserRequest struct {
	ID string `json:"id"`
}
type fetchUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Err      string
}

// FetchUserEndpoint
func FetchUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(fetchUserRequest)
		v, err := svc.FetchUser(req.ID)
		if err != nil {
			return fetchUserResponse{v.ID, v.Username, err.Error()}, nil
		}
		return fetchUserResponse{v.ID, v.Username, ""}, nil
	}
}

func DeleteUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(fetchUserRequest)
		v, err := svc.DeleteUser(req.ID)
		if err != nil {
			return fetchUserResponse{v.ID, v.Username, err.Error()}, nil
		}
		return fetchUserResponse{v.ID, v.Username, ""}, nil
	}
}

// DecodeFetchUserRequest
func DecodeFetchUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req fetchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// EncodeFetchUserResponse
func EncodeFetchUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type createUserRequest struct {
	Username string `json:"username"`
	RegionID int    `json:"region_id"`
}

// CreateUserEndpoint
func CreateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		v, err := svc.CreateUser(req.Username, req.RegionID)
		if err != nil {
			return fetchUserResponse{v.ID, v.Username, err.Error()}, nil
		}
		return fetchUserResponse{v.ID, v.Username, ""}, nil
	}
}

// DecodeCreateUserRequest
func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

type updateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	RegionID int    `json:"region_id"`
}

func UpdateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		v, err := svc.UpdateUser(req.ID, req.Username)
		if err != nil {
			return fetchUserResponse{v.ID, v.Username, err.Error()}, nil
		}
		return fetchUserResponse{v.ID, v.Username, ""}, nil
	}
}

func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
