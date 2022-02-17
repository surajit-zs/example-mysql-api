package cat

import (
	"net/http"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"mysql-api/models"
	"mysql-api/services"
)

type Handler struct {
	Services services.Services
}

func (h Handler) Get(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.Services.Get(ctx)

	if err != nil {
		return nil, errors.DB{Err: errors.Error("Internal DB error")}
	}

	res := models.Response{
		Cat:        resp,
		Massage:    "Success",
		StatusCode: http.StatusOK,
	}

	return res, nil
}

func (h Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var c models.Cat

	if err := ctx.Bind(&c); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if c.ID == "0" {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.Services.Create(ctx, c)
	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	res := models.Response{
		Cat:        resp,
		Massage:    "created successful",
		StatusCode: http.StatusCreated,
	}

	return res, nil
}
