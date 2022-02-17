package cat

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"mysql-api/models"
	"mysql-api/store"
)

type Services struct {
	Store store.Store
}

func New(s store.Store) Services {
	return Services{Store: s}
}

func (s Services) Get(ctx *gofr.Context) ([]models.Cat, error) {
	res, err := s.Store.Get(ctx)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return res, nil
}
func (s Services) Create(ctx *gofr.Context, c models.Cat) (models.Cat, error) {
	if !check(c) {
		return models.Cat{}, errors.Error("validation error")
	}

	res, err := s.Store.Create(ctx, c)
	if err != nil {
		return models.Cat{}, errors.DB{Err: err}
	}

	return res, nil
}

func check(c models.Cat) bool {
	id, err := strconv.Atoi(c.ID)

	if err != nil {
		return false
	}

	if id <= 0 {
		return false
	}

	if c.Name == "" || c.Age <= 0 {
		return false
	}

	return true
}
