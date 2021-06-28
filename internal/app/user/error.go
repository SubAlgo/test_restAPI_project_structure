package user

import (
	"errors"
	"net/http"
)

var (
	ErrNameRequire     = errors.New("name require")
	ErrLastnameRequire = errors.New("lastname require")
	ErrEmailRequire    = errors.New("email require")
	ErrMethodNotAllow  = errors.New("method not allow")
	ErrCustomerUpdate  = errors.New("error: customer update data")
	ErrDeleteCustomer  = errors.New("error: delete customer data")
	ErrSalaryRequire   = errors.New("error: salary require")
)

func ErrorToStatusCode(err error) int {
	switch err {
	case ErrNameRequire, ErrLastnameRequire, ErrEmailRequire, ErrSalaryRequire:
		return http.StatusBadRequest
	case ErrMethodNotAllow:
		return http.StatusMethodNotAllowed
	case ErrCustomerUpdate, ErrDeleteCustomer:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func ErrorToMessage(err error) string {
	switch err {
	case ErrMethodNotAllow:
		return "Method not allowed"
	case ErrSalaryRequire:
		return "please input salary"
	case ErrCustomerUpdate:
		return "update customer data not success"
	case ErrDeleteCustomer:
		return "delete customer data not success"
	default:
		return "Internal server error"
	}
}
