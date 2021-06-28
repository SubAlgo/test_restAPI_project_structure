package user

import (
	"context"
	"database/sql"
	mUser "github.com/subalogo/bookRestAPI/internal/model/user"
	iUser "github.com/subalogo/bookRestAPI/internal/repo/user"
	"strconv"
)

type Customer iUser.Customer

func (c *Customer) Create(ctx context.Context) (res iUser.ResponseMessage, err error) {

	if c.Name == "" {
		return res, ErrNameRequire
	}

	if c.Lastname == "" {
		return res, ErrLastnameRequire
	}

	if c.Email == "" {
		return res, ErrEmailRequire
	}

	var salary = 0.0

	var uStruct = iUser.UserStruct{
		Role:     "customer",
		Name:     c.Name,
		Lastname: c.Lastname,
		Email:    c.Email,
		Salary:   salary,
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

func (c *Customer) Get(ctx context.Context, idStr string) (res iUser.UserStruct, err error) {
	var id int
	id, err = strconv.Atoi(idStr)
	res, err = mUser.Get(ctx, id)
	if err == sql.ErrNoRows {
		return iUser.UserStruct{}, nil
	}
	if err != nil {
		return iUser.UserStruct{}, err
	}
	return res, err
}

func (c *Customer) List(ctx context.Context) (res iUser.UserList, err error) {
	res, err = mUser.GetCustomerList(ctx)
	if err != nil {
		return iUser.UserList{}, err
	}
	return res, err
}

func (c *Customer) Update(ctx context.Context) (res iUser.ResponseMessage, err error) {

	if c.Name == "" {
		return res, ErrNameRequire
	}

	if c.Lastname == "" {
		return res, ErrLastnameRequire
	}

	if c.Email == "" {
		return res, ErrEmailRequire
	}

	var uStruct = iUser.UserStruct{
		ID:       c.ID,
		Name:     c.Name,
		Lastname: c.Lastname,
		Email:    c.Email,
		Salary:   0.0,
		Role:     "customer",
	}

	err = mUser.Update(ctx, &uStruct)
	if err != nil {
		return res, ErrCustomerUpdate
	}

	res.Message = "Update customer data success"

	return res, err
}

func (c *Customer) Delete(ctx context.Context, id int) (res iUser.ResponseMessage, err error) {

	err = mUser.Del(ctx, id)

	if err != nil {
		return iUser.ResponseMessage{}, ErrDeleteCustomer
	}

	res.Message = "delete customer data success"
	return res, err
}
