package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/huzaifamk/Go-Clean-Arch-Project-1/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type BookHandler struct {
	Bookservice models.BookService
}

func NewBookHandler(e *echo.Echo, se models.BookService) {
	handler := &BookHandler{
		Bookservice: se,
	}
	e.GET("/book/:id", handler.GetByID)
	e.POST("/books/add", handler.Store)

}

func (a *BookHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.Bookservice.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *models.Book) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *BookHandler) Store(c echo.Context) (err error) {

	var book models.Book
	err = c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&book); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.Bookservice.Store(ctx, &book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, book)
}
