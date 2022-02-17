package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"mysql-api/http/cat"
	services "mysql-api/services/cat"
	store "mysql-api/store/cat"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	st := store.New()
	s := services.New(st)
	h := cat.Handler{Services: s}

	app.GET("/cat", h.Get)
	app.POST("/cat", h.Create)

	app.Start()
}
