package user

import (
	"context"
	mUser "github.com/subalogo/bookRestAPI/internal/model/user"
	iUser "github.com/subalogo/bookRestAPI/internal/repo/user"
)

func List(ctx context.Context) (res iUser.UserList, err error) {
	res, err = mUser.GetList(ctx)
	if err != nil {
		return iUser.UserList{}, err
	}
	return res, err
}
