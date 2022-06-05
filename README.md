# GO Library REST API
A RESTful API example for simple todo application with Go
# Instalation and Run
```
# Download this project
go get github.com/gxrlxv/library-rest-api
```
Before running API server, you should set the database config with yours or set the your database config with my values on config.go
```go
func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
```
Build and Run
```
cd library-rest-api
go build
./library-rest-api
```
# API
#### /api/users/
- GET all users
#### /api/users/:user_id
- GET get user by id
- PUT update user
- DELETE delete user
#### /api/users/sign-up
- POST sign up users
#### /api/users/sign-in
- POST sign up users
#### /api/authors
- POST create authors
- GET all authors
#### /api/authors/:author_id
- GET get author by id
- PUT update user
- DELETE delete user 
#### /api/books
- POST create books
- GET all books
#### /api/books/:book_id
- GET get book by id
- PUT update book
- DELETE delete book 
#### /api/books/:book_id/take
- PUT take a book
- DELETE undo a book 
