package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	handler "github.com/HMasataka/onion/application/api"
	"github.com/HMasataka/onion/application/api/router"
	"github.com/HMasataka/onion/application/usecase"
	"github.com/HMasataka/onion/infrastructure/repository"
	"github.com/HMasataka/onion/transaction"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// user:password@tcp(host:port)/dbname
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		"user", "password", "localhost:3306", "db",
	))
	if err != nil {
		panic(err)
	}

	connectionProvider := transaction.NewConnectionProvider(db)
	transactor := transaction.NewTransactor(connectionProvider)

	userRepository := repository.NewUserRepository(connectionProvider)
	userUseCase := usecase.NewUserUseCase(transactor, userRepository)

	r.Group(func(r chi.Router) {
		r.Get("/users", router.NewHandler(handler.NewFindUserHandler(userUseCase)))
	})

	http.ListenAndServe(":3000", r)
}
