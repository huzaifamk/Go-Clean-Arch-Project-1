package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_bookHandler "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/controller"
	_bookHandlerMiddleware "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/controller/middleware"
	_bookRepo "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/repository/mysql"
	_bookService "github.com/huzaifamk/Go-Clean-Arch-Project-1/books/service"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Server running successfully on port 9090!")
	}
}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _bookHandlerMiddleware.InitMiddleware()
	e.Use(middL.Logger())
	e.Use(middL.CORS)
	br := _bookRepo.NewMysqlBookRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	bs := _bookService.NewBookService(br, timeoutContext)
	_bookHandler.NewBookHandler(e, bs)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
