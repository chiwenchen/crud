package main

import (
	"encoding/json"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func (mw loggingMiddleware) FetchUser(id string) (output *User, err error) {
	defer func(begin time.Time) {

		usr, err := json.Marshal(output)

		mw.logger.Log(
			"method", "find_user",
			"input", id,
			"output", usr,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.FetchUser(id)
	return
}

func (mw loggingMiddleware) CreateUser(username string, regionID int) (output *User, err error) {

	input := struct {
		username string
		regionID int
	}{username, regionID}

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.CreateUser(input.username, input.regionID)
	return
}

func (mw loggingMiddleware) UpdateUser(id string, username string) (output *User, err error) {

	input := struct {
		id       string
		username string
	}{id, username}

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.UpdateUser(input.id, input.username)
	return
}

func (mw loggingMiddleware) DeleteUser(id string) (output *User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.DeleteUser(id)
	return
}
