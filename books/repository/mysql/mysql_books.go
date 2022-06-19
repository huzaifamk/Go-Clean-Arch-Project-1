package mysql

import (
	"context"
	"database/sql"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models"

	"github.com/sirupsen/logrus"
)

type mysqlBookRepository struct {
	Conn *sql.DB
}

func NewMysqlBookRepository(Conn *sql.DB) models.BookRepository {
	return &mysqlBookRepository{Conn}
}

func (m *mysqlBookRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []models.Book, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]models.Book, 0)
	for rows.Next() {
		t := models.Book{}
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.Author,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlBookRepository) GetByID(ctx context.Context, id int64) (res models.Book, err error) {
	query := `SELECT id, title, content, author_name, created_at
  						FROM books WHERE id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return models.Book{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, models.ErrNotFound
	}

	return
}

func (m *mysqlBookRepository) Store(ctx context.Context, a *models.Book) (err error) {
	query := `INSERT  books SET id=? , title=? , content=? , author_name=?, created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, a.ID, a.Title, a.Content, a.Author, a.CreatedAt)
	if err != nil {
		return
	}

	return

}
