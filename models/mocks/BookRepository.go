package mocks

import (
	context "context"

	models "github.com/huzaifamk/Go-Clean-Arch-Project-1/models"
	mock "github.com/stretchr/testify/mock"
)

type BookRepository struct {
	mock.Mock
}

func (_m *BookRepository) GetByID(ctx context.Context, id int64) (models.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.Book); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *BookRepository) Store(_a0 context.Context, _a1 *models.Book) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Book) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
