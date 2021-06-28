package main // import "github.com/subalogo/bookRestAPI"
import (
	"database/sql"
	bookController "github.com/subalogo/bookRestAPI/internal/controller/book"
	userController "github.com/subalogo/bookRestAPI/internal/controller/user"
	"github.com/subalogo/bookRestAPI/internal/pkg/configDB"
	"github.com/subalogo/bookRestAPI/internal/pkg/dbctx"
	"log"
	"net/http"
)

func main() {
	var err error

	var db *sql.DB
	db = configDB.Postgres()
	defer db.Close()
	if db.Ping() != nil {
		log.Fatal("connect main db fail")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.NotFoundHandler())
	//mux.Handle("/book", book.Handler())
	mux.Handle("/api/book", http.StripPrefix("/api", bookController.Handler()))
	mux.Handle("/api/user/", http.StripPrefix("/api/user", userController.Handler()))
	h := chain(
		dbctx.Middleware(db),
	)(mux)

	log.Println("staring server on :8000")
	if err = http.ListenAndServe(":8000", h); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func chain(hs ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}
