package book

import "context"

// IBook : interface book
type IBook interface {
	Create(context.Context) (ResponseMessage, error)
	List(context.Context) (BookList, error)
	Get(context.Context, string) (Book, error)
	Update(context.Context) (ResponseMessage, error)
	Delete(context.Context) (ResponseMessage, error)
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type BookList struct {
	Count int    `json:"count"`
	List  []Book `json:"list"`
}
