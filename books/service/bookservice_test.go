package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	service "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/service"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models/mocks"
)

func TestGetByID(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := models.Book{
		Title:   "Hello",
		Content: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockBookRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockBook, nil).Once()
		u := service.NewBookService(mockBookRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockBook.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockBookRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockBookRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(models.Book{}, errors.New("Unexpected")).Once()

		u := service.NewBookService(mockBookRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockBook.ID)

		assert.Error(t, err)
		assert.Equal(t, models.Book{}, a)

		mockBookRepo.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBook := models.Book{
		Title:   "Hello",
		Content: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempmockBook := mockBook
		tempmockBook.ID = 0
		mockBookRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(models.Book{}, models.ErrNotFound).Once()
		mockBookRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Book")).Return(nil).Once()

		u := service.NewBookService(mockBookRepo, time.Second*2)

		err := u.Store(context.TODO(), &tempmockBook)

		assert.NoError(t, err)
		assert.Equal(t, mockBook.Title, tempmockBook.Title)
		mockBookRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingBook := mockBook
		mockBookRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(existingBook, nil).Once()

		u := service.NewBookService(mockBookRepo, time.Second*2)

		err := u.Store(context.TODO(), &mockBook)

		assert.Error(t, err)
		mockBookRepo.AssertExpectations(t)
	})

}
