# Golang Clean Architecture Sample Project

This project is still in development, and I will add more features in the future. It is made by following CLean Architecture rules by Uncle Bob, and I also added some of my own rules to make it more readable and maintainable. This is written in Golang, and I used Echo Framework as the web framework, and MySQL as the database. I also added some middlewares to make it more secure and more readable. I also added some unit tests to make sure that the code is working as expected.

## How to run

1. Clone this repository
2. Create a database named `books.sql` in your MySQL database [Name is changeable]
3. Import the `books.sql` file in the database
4. Change the database configuration in `config.json` file
5. Run the project using `go run main.go` or refer to Makefile for more commands

## API Documentation

### Get All Books

#### Request

`GET /books`

    curl -i -H 'Accept: application/json' http://localhost:8080/books

#### Response

```json
    HTTP/1.1 200 OK
    Date: Sun, 20 Jun 2021 15:00:00 GMT
    Content-Type: application/json
    Content-Length: 2
    
        [
            {
                "id": 1,
                "title": "Book Title",
                "author": "Book Author"
            }
        ]

```

### Get Book by ID

#### Request

`GET /books/:id`

    curl -i -H 'Accept: application/json' http://localhost:8080/books/1

#### Response

    HTTP/1.1 200 OK
    Date: Sun, 20 Jun 2021 15:00:00 GMT
    Content-Type: application/json
    Content-Length: 2

    []

### Create a Book

#### Request

`POST /books`

    curl -i -H 'Accept: application/json' -d 'title=Book Title&author=Book Author' http://localhost:8080/books

#### Response

    HTTP/1.1 201 Created
    Date: Sun, 20 Jun 2021 15:00:00 GMT
    Content-Type: application/json
    Content-Length: 2

    []

### Update a Book

#### Request

`PUT /books/:id`

    curl -i -H 'Accept: application/json' -d 'title=Book Title&author=Book Author' http://localhost:8080/books/1

#### Response

    HTTP/1.1 200 OK
    Date: Sun, 20 Jun 2021 15:00:00 GMT
    Content-Type: application/json
    Content-Length: 2

    []

### Delete a Book

#### Request

`DELETE /books/:id`

    curl -i -H 'Accept: application/json' -X DELETE http://localhost:8080/books/1

#### Response

    HTTP/1.1 204 No Content
    Date: Sun, 20 Jun 2021 15:00:00 GMT


## License

MIT

## Author

[Huzaifa M.]