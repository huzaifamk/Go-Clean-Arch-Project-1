package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	bookController "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/controller"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models"
	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models/mocks"
)

// func TestFetch(t *testing.T) {
// 	var mockBook models.Book
// 	err := faker.FakeData(&mockBook)
// 	assert.NoError(t, err)
// 	mockService := new(mocks.BookService)
// 	mockListBook := make([]models.Book, 0)
// 	mockListBook = append(mockListBook, mockBook)
// 	num := 1
// 	cursor := "2"
// 	mockService.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListBook, "10", nil)

// 	e := echo.New()
// 	req, err := http.NewRequest(echo.GET, "/book?id=1&cursor="+cursor, strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := bookController.BookHandler{
// 		Bookservice: mockService,
// 	}
// 	err = handler.GetByID(c)
// 	require.NoError(t, err)

// 	responseCursor := rec.Header().Get("X-Cursor")
// 	assert.Equal(t, "10", responseCursor)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	mockService.AssertExpectations(t)
// }

// func TestFetchError(t *testing.T) {
// 	mockService := new(mocks.BookService)
// 	num := 1
// 	cursor := "2"
// 	mockService.On("Fetch", mock.Anything, cursor, int64(num)).Return(nil, "", models.ErrInternalServerError)

// 	e := echo.New()
// 	req, err := http.NewRequest(echo.GET, "/book?id=1&cursor="+cursor, strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := bookController.BookHandler{
// 		Bookservice: mockService,
// 	}
// 	err = handler.GetByID(c)
// 	require.NoError(t, err)

// 	responseCursor := rec.Header().Get("X-Cursor")
// 	assert.Equal(t, "", responseCursor)
// 	assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	mockService.AssertExpectations(t)
// }

func TestGetByID(t *testing.T) {
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockService := new(mocks.BookService)

	num := int(mockBook.ID)

	mockService.On("GetByID", mock.Anything, int64(num)).Return(mockBook, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/book/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("book/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := bookController.BookHandler{
		Bookservice: mockService,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockBook := models.Book{
		ID:        123,
		Title:     "Title",
		Content:   "Content",
		Author:    "Author",
		CreatedAt: time.Now(),
	}

	tempmockBook := mockBook
	tempmockBook.ID = 0
	mockService := new(mocks.BookService)

	j, err := json.Marshal(tempmockBook)
	assert.NoError(t, err)

	mockService.On("Store", mock.Anything, mock.AnythingOfType("*models.Book")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/book", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/book")

	handler := bookController.BookHandler{
		Bookservice: mockService,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}
