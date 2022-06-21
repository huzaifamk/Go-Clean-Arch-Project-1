package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	bookMysqlRepo "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/repository/mysql"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_name", "created_at"}).
		AddRow(1, "title 1", "Content 1", 1, time.Now(), time.Now())

	query := "SELECT title,content, author_name, created_at FROM Book WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := bookMysqlRepo.NewMysqlBookRepository(db)

	num := int64(5)
	anBook, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anBook)
}

func TestStore(t *testing.T) {
	now := time.Now()
	ar := &models.Book{
		ID:        123,
		Title:     "Judul",
		Content:   "Content",
		Author:    "Iman Tumorang",
		CreatedAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT  Book SET id=\\? , title=\\? , content=\\? , author_name=\\?, created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.ID, ar.Title, ar.Content, ar.Author, ar.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := bookMysqlRepo.NewMysqlBookRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}
