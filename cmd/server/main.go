package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"

	handler "github.com/HMasataka/onion/application/api"
	"github.com/HMasataka/onion/application/api/router"
	"github.com/HMasataka/onion/application/usecase"
	"github.com/HMasataka/onion/infrastructure/repository"
	"github.com/HMasataka/onion/transaction"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	c := mysql.Config{
		DBName:    "db",
		User:      "user",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
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
