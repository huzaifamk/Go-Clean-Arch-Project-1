package models

import (
	"context"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author_name"`
	CreatedAt time.Time `json:"created_at"`
}

type BookService interface {
	GetByID(ctx context.Context, id int64) (Book, error)
	Store(context.Context, *Book) (error)
}

type BookRepository interface {
	GetByID(ctx context.Context, id int64) (Book, error)
	Store(ctx context.Context, a *Book) error
}

