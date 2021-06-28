package user

import (
	"context"
	"database/sql"
	"github.com/subalogo/bookRestAPI/internal/pkg/dbctx"
	iUser "github.com/subalogo/bookRestAPI/internal/repo/user"
)

func Register(ctx context.Context, u *iUser.UserStruct) (id int64, err error) {
	err = dbctx.QueryRow(ctx, `
		INSERT INTO 
			users
			(name, lastname, email, salary, role)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING
			id
	`, u.Name, u.Lastname, u.Email, u.Salary, u.Role).Scan(&id)
	return id, err
}

func GetList(ctx context.Context) (res iUser.UserList, err error) {
	var rows *sql.Rows

	rows, err = dbctx.Query(ctx, `
		SELECT 
			id, name, lastname, email, role, salary
		FROM
			users
	`)

	if err == sql.ErrNoRows {
		return iUser.UserList{}, nil
	}
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var x iUser.UserStruct
		if err = rows.Scan(&x.ID, &x.Name, &x.Lastname, &x.Email, &x.Role, &x.Salary); err != nil {
			return iUser.UserList{}, err
		}
		res.List = append(res.List, x)
	}

	res.Count = len(res.List)
	return res, err
}

func GetCustomerList(ctx context.Context) (res iUser.UserList, err error) {
	var rows *sql.Rows

	rows, err = dbctx.Query(ctx, `
		SELECT 
			id, name, lastname, email, role, salary
		FROM
			users
		WHERE
			role = 'customer'
		`)

	if err == sql.ErrNoRows {
		return iUser.UserList{}, nil
	}
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var x iUser.UserStruct
		if err = rows.Scan(&x.ID, &x.Name, &x.Lastname, &x.Email, &x.Role, &x.Salary); err != nil {
			return iUser.UserList{}, err
		}
		res.List = append(res.List, x)
	}

	res.Count = len(res.List)
	return res, err
}

func GetListEmployee(ctx context.Context) (res iUser.UserList, err error) {

	var rows *sql.Rows

	rows, err = dbctx.Query(ctx, `
		SELECT 
			id, name, lastname, email, role, salary
		FROM
			users
		WHERE
			role <> 'customer'
		`)

	if err == sql.ErrNoRows {
		return iUser.UserList{}, nil
	}
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var x iUser.UserStruct
		if err = rows.Scan(&x.ID, &x.Name, &x.Lastname, &x.Email, &x.Role, &x.Salary); err != nil {
			return iUser.UserList{}, err
		}
		res.List = append(res.List, x)
	}

	res.Count = len(res.List)
	return res, err
}

func Get(ctx context.Context, id int) (res iUser.UserStruct, err error) {
	err = dbctx.QueryRow(ctx, `
		SELECT 
			id, name, lastname, email, salary, role
		FROM
			users
		WHERE
			id = $1
	`, id).Scan(&res.ID, &res.Name, &res.Lastname, &res.Email, &res.Salary, &res.Role)
	return res, err
}

func Update(ctx context.Context, u *iUser.UserStruct) (err error) {
	_, err = dbctx.Exec(ctx, `
		UPDATE 
			users
		SET
			name = $2,
			lastname = $3,
			email = $4,
			salary = $5,
			role = $6
		WHERE
			id = $1
	`, u.ID, u.Name, u.Lastname, u.Email, u.Salary, u.Role)
	return err
}

func Del(ctx context.Context, id int) (err error) {
	_, err = dbctx.Exec(ctx, `
		DELETE FROM users WHERE id = $1
	`, id)
	return err
}
