package user

import "context"

type IUser interface {
	Create(context.Context) (ResponseMessage, error)
	Get(context.Context, string) (UserStruct, error)
	List(context.Context) (UserList, error)
	Update(context.Context) (ResponseMessage, error)
	Delete(context.Context, int) (ResponseMessage, error)
}

type UserStruct struct {
	ID       int     `json:"id"`
	Role     string  `json:"role"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Email    string  `json:"email"`
	Salary   float64 `json:"salary"`
}

type ResponseMessage struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type UserList struct {
	Count int          `json:"count"`
	List  []UserStruct `json:"list"`
}
