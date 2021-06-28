package user

import (
	"fmt"
	appUser "github.com/subalogo/bookRestAPI/internal/app/user"
	"github.com/subalogo/bookRestAPI/internal/pkg/transport"
	"net/http"
)

var t = transport.HTTP{
	ErrorToStatusCode: appUser.ErrorToStatusCode,
	ErrorToMessage:    appUser.ErrorToMessage,
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("/list", userHandler)
	mux.HandleFunc("/customer", customerHandler)
	mux.HandleFunc("/employee", employeeHandler)

	return mux
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		err error
		res interface{}
	)

	switch r.Method {
	case http.MethodGet:
		res, err = appUser.List(ctx)
	default:
		err = appUser.ErrMethodNotAllow
	}

	t.EncodeResult(w, res, err)
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		err error
		res interface{}
		c   appUser.Customer
	)

	if err = t.DecodeRequest(w, r, &c); err != nil {
		t.EncodeResult(w, r, err)
	}

	switch r.Method {
	case http.MethodPost:
		res, err = c.Create(ctx)
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		fmt.Println("idStr: ", idStr)
		switch idStr {
		case "":
			res, err = c.List(ctx)
		default:
			res, err = c.Get(ctx, idStr)
		}
	case http.MethodPut:
		res, err = c.Update(ctx)
	case http.MethodDelete:
		res, err = c.Delete(ctx, c.ID)
	default:
		err = appUser.ErrMethodNotAllow
	}

	t.EncodeResult(w, res, err)
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		err error
		res interface{}
		e   appUser.Employee
	)

	if err = t.DecodeRequest(w, r, &e); err != nil {
		t.EncodeResult(w, r, err)
	}

	switch r.Method {
	case http.MethodPost:
		res, err = e.Create(ctx)
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		switch idStr {
		case "":
			res, err = e.List(ctx)
		default:
			res, err = e.Get(ctx, idStr)
		}
	case http.MethodPut:
		res, err = e.Update(ctx)
	case http.MethodDelete:
		res, err = e.Delete(ctx, e.ID)
	default:
		err = appUser.ErrMethodNotAllow
	}

	t.EncodeResult(w, res, err)
}
