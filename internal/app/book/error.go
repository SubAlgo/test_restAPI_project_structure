package book

import (
	"errors"
	"net/http"
)

var (
	ErrNameRequire = errors.New("")
)

func ErrorToStatusCode(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

func ErrorToMessage(err error) string {
	switch err {
	default:
		return "Internal server error"
	}
}
