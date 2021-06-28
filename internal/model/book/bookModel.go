package book

import (
	"context"
	"database/sql"
	"github.com/subalogo/bookRestAPI/internal/pkg/dbctx"
	iBook "github.com/subalogo/bookRestAPI/internal/repo/book"
)

func Create(ctx context.Context, b *iBook.Book) error {
	_, err := dbctx.Exec(ctx, `
		INSERT INTO 
			book
			(title, author)
		VALUES
			($1, $2)
	`, b.Title, b.Author)
	if err != nil {
		return err
	}
	return nil
}

func GetList(ctx context.Context) (res iBook.BookList, err error) {
	var rows *sql.Rows
	rows, err = dbctx.Query(ctx, `
		SELECT 
			id, title, author
		FROM
			book
	`)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var x iBook.Book
		rows.Scan(&x.ID, &x.Title, &x.Author)
		res.List = append(res.List, x)
	}
	return res, err
}

func Get(ctx context.Context, id int) (iBook.Book, error) {
	var b iBook.Book
	var err error

	err = dbctx.QueryRow(ctx, `
		SELECT 
			id, title, author
		FROM
			book
		WHERE
			id = $1
	`, id).Scan(&b.ID, &b.Title, &b.Author)
	return b, err
}

func Update(ctx context.Context, b *iBook.Book) (err error) {
	_, err = dbctx.Exec(ctx, `
		UPDATE 
			book
		SET
			title = $2,
			author = $3
		WHERE
			id = $1
	`, b.ID, b.Title, b.Author)
	return err
}

func Del(ctx context.Context, id int) (err error) {
	_, err = dbctx.Exec(ctx, `
		DELETE FROM book WHERE id = $1
	`, id)
	return err
}
