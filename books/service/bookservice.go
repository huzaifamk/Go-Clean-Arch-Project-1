package service

import (
	"context"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models"
	"time"
)

type ResponseError struct {
	Message string `json:"message"`
}

type bookService struct {
	bookRepo       models.BookRepository
	contextTimeout time.Duration
}

func NewBookService(b models.BookRepository, timeout time.Duration) models.BookService {

	return &bookService{
		bookRepo:       b,
		contextTimeout: timeout,
	}
}

func (a *bookService) GetByID(c context.Context, id int64) (res models.Book, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.bookRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *bookService) Store(c context.Context, m *models.Book) (err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedBook, _ := a.GetByID(ctx, m.ID)
	if existedBook != (models.Book{}) {
		return models.ErrConflict
	}

	err = a.bookRepo.Store(ctx, m)
	return
}
