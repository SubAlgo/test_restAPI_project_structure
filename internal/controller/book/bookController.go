package book

import (
	"context"
	appBook "github.com/subalogo/bookRestAPI/internal/app/book"
	"github.com/subalogo/bookRestAPI/internal/pkg/transport"
	interfaceBook "github.com/subalogo/bookRestAPI/internal/repo/book"
	"net/http"
)

var t = transport.HTTP{
	ErrorToStatusCode: appBook.ErrorToStatusCode,
	ErrorToMessage:    appBook.ErrorToMessage,
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("/book", bookHandler)

	return mux
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	var (
		b   appBook.Book
		res interface{}
		ctx context.Context
		err error
		id  string
	)

	ctx = r.Context()

	if err = t.DecodeRequest(w, r, &b); err != nil {
		t.EncodeResult(w, r, err)
	}

	var ib = interfaceBook.IBook(&b)

	switch r.Method {
	case http.MethodGet:
		id = r.URL.Query().Get("id")
		switch id {
		case "":
			res, err = ib.List(ctx)
		default:
			res, err = ib.Get(ctx, id)
		}
	case http.MethodPost:
		res, err = ib.Create(ctx)
	case http.MethodPut:
		res, err = ib.Update(ctx)
	case http.MethodDelete:
		res, err = b.Delete(ctx)
	}

	t.EncodeResult(w, res, err)
}
