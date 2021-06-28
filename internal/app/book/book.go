package book

import (
	"context"
	"errors"
	mBook "github.com/subalogo/bookRestAPI/internal/model/book"
	iBook "github.com/subalogo/bookRestAPI/internal/repo/book"
	"strconv"
)

func (b *Book) Create(ctx context.Context) (res iBook.ResponseMessage, err error) {
	if b.Title == "" {
		return res, errors.New("title invalid")
	}

	if b.Author == "" {
		return res, errors.New("author invalid")
	}

	var nBook = iBook.Book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
	}

	err = mBook.Create(ctx, &nBook)
	res.Message = "บันทึกข้อมูลสำเร็จ"
	return res, nil
}

func (b *Book) List(ctx context.Context) (res iBook.BookList, err error) {
	res, err = mBook.GetList(ctx)
	if err != nil {
		return res, err
	}
	res.Count = len(res.List)
	return res, err
}

func (b *Book) Get(ctx context.Context, id string) (res iBook.Book, err error) {

	var nID int
	nID, err = strconv.Atoi(id)
	if err != nil {
		return res, err
	}
	res, err = mBook.Get(ctx, nID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b *Book) Update(ctx context.Context) (res iBook.ResponseMessage, err error) {
	if b.Title == "" {
		return res, errors.New("title invalid")
	}

	if b.Author == "" {
		return res, errors.New("title invalid")
	}
	var nBook = iBook.Book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
	}
	err = mBook.Update(ctx, &nBook)
	if err != nil {
		return res, err
	}
	res.Message = "update data success"
	return res, nil
}

func (b *Book) Delete(ctx context.Context) (res iBook.ResponseMessage, err error) {
	if err = mBook.Del(ctx, b.ID); err != nil {
		return res, err
	}
	res.Message = "Delete data success"
	return res, err
}
