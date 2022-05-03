package main

import (
	"mysql-api/middleware"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"mysql-api/http/cat"
	services "mysql-api/services/cat"
	store "mysql-api/store/cat"
)

func main() {
	app := gofr.New()

	app.Server.UseMiddleware(middleware.Handler)

	st := store.New()
	s := services.New(st)
	h := cat.Handler{Services: s}

	app.GET("/cat", h.Get)
	app.GET("/cat/{id}", h.GetByID)
	app.POST("/cat", h.Create)
	app.PUT("/cat/{id}", h.Update)
	app.DELETE("/cat/{id}", h.Delete)

	app.Start()
}
