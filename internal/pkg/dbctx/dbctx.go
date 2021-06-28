package dbctx

import (
	"context"
	"database/sql"
	"net/http"
)

// ctx key
type (
	dbKey    struct{}
	queryKey struct{}
)

// create query type interface
type query interface {
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

// NewContext create new context
func NewContext(parent context.Context, db *sql.DB) context.Context {
	ctx := context.WithValue(parent, dbKey{}, db)
	ctx = context.WithValue(ctx, queryKey{}, db)
	return ctx
}

func Middleware(db *sql.DB) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = NewContext(ctx, db)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func getDB(ctx context.Context) *sql.DB {
	return ctx.Value(dbKey{}).(*sql.DB)
}

func GetDB(ctx context.Context) *sql.DB {
	return ctx.Value(dbKey{}).(*sql.DB)
}

func getQuery(ctx context.Context) query {
	return ctx.Value(queryKey{}).(query)
}

type Txfunc func(ctx context.Context) error

func RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
	db := getDB(ctx)
	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	// retry
	//db.PrepareContext()

	var tx *sql.Tx
	var err error

	tx, err = db.BeginTx(ctx, opt)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txctx := context.WithValue(ctx, queryKey{}, tx)
	// type Txfunc func(ctx context.Context) error
	err = f(txctx)
	if err != nil {
		// retryable ?
		return err

	}
	// error และ หมดโควต้า retry
	return tx.Commit()
	// ------
}

func QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return getQuery(ctx).QueryRowContext(ctx, query, args...)
}

func Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return getQuery(ctx).QueryContext(ctx, query, args...)
}

func Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return getQuery(ctx).ExecContext(ctx, query, args...)
}
