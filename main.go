package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

var errEmpty = errors.New("empty string")
var logger = log.NewLogfmtLogger(os.Stderr)

func outerloggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}

func main() {
	svc := userService{}

	fetchUser := endpoint.Chain(
		outerloggingMiddleware(log.With(logger, "method", "fetch_user")),
	)(FetchUserEndpoint(svc))

	fetchUserHandler := httptransport.NewServer(
		fetchUser,
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
	http.ListenAndServe(":8090", nil)

	// r := mux.NewRouter()
	// e := MakeServerEndpoints(s)

	// r.Methods("GET").Path("/find_user").Handler(httptransport.NewServer(
	// 	FetchUserEndpoint(svc),
	// 	DecodeFetchUserRequest,
	// 	EncodeFetchUserResponse,
	// ))

	// return r
}
