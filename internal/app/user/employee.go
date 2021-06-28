package user

import (
	"context"
	"database/sql"
	mUser "github.com/subalogo/bookRestAPI/internal/model/user"
	iUser "github.com/subalogo/bookRestAPI/internal/repo/user"
	"strconv"
	"strings"
)

type Employee iUser.EmployeeStruct

func (e *Employee) Create(ctx context.Context) (res iUser.ResponseMessage, err error) {
	if e.Name == "" {
		return res, ErrNameRequire
	}

	if e.Lastname == "" {
		return res, ErrLastnameRequire
	}

	if e.Email == "" {
		return res, ErrEmailRequire
	}

	if strings.TrimSpace(e.Role) == "" {
		e.Role = "employee"
	}

	if e.Salary <= 0 {
		return iUser.ResponseMessage{}, ErrSalaryRequire
	}

	var uStruct = iUser.UserStruct{
		Role:     e.Role,
		Name:     e.Name,
		Lastname: e.Lastname,
		Email:    e.Email,
		Salary:   e.Salary,
	}
	var userID int64
	userID, err = mUser.Register(ctx, &uStruct)
	if err != nil {
		return res, err
	}

	res.ID = userID
	res.Message = "register success"
	return res, err
}

func (e *Employee) Get(ctx context.Context, idStr string) (res iUser.UserStruct, err error) {
	var id int
	id, err = strconv.Atoi(idStr)
	res, err = mUser.Get(ctx, id)
	if err == sql.ErrNoRows {
		return res, nil
	}
	if err != nil {
		return iUser.UserStruct{}, err
	}
	return res, err
}

func (e *Employee) List(ctx context.Context) (res iUser.UserList, err error) {
	res, err = mUser.GetListEmployee(ctx)
	if err != nil {
		return iUser.UserList{}, err
	}
	return res, err
}

func (e *Employee) Update(ctx context.Context) (res iUser.ResponseMessage, err error) {

	if e.Name == "" {
		return res, ErrNameRequire
	}

	if e.Lastname == "" {
		return res, ErrLastnameRequire
	}

	if e.Email == "" {
		return res, ErrEmailRequire
	}

	if e.Salary <= 0 {
		return iUser.ResponseMessage{}, ErrSalaryRequire
	}

	if strings.TrimSpace(e.Role) == "" {
		e.Role = "employee"
	}

	var uStruct = iUser.UserStruct{
		ID:       e.ID,
		Name:     e.Name,
		Lastname: e.Lastname,
		Email:    e.Email,
		Salary:   e.Salary,
		Role:     e.Role,
	}

	err = mUser.Update(ctx, &uStruct)
	if err != nil {
		return res, ErrCustomerUpdate
	}

	res.Message = "Update customer data success"

	return res, err
}

func (e *Employee) Delete(ctx context.Context, id int) (res iUser.ResponseMessage, err error) {
	err = mUser.Del(ctx, id)

	if err != nil {
		return iUser.ResponseMessage{}, ErrDeleteCustomer
	}

	res.Message = "delete customer data success"
	return res, err
}
