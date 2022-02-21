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

func (s Services) GetByID(ctx *gofr.Context, id string) (models.Cat, error) {
	if !idCheck(id) {
		return models.Cat{}, errors.Error("validation error")
	}

	res, err := s.Store.GetByID(ctx, id)
	if err != nil {
		return models.Cat{}, errors.DB{Err: err}
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

func (s Services) Update(ctx *gofr.Context, c models.Cat) (models.Cat, error) {
	if !check(c) {
		return models.Cat{}, errors.Error("validation error")
	}

	res, err := s.Store.Update(ctx, c)
	if err != nil {
		return models.Cat{}, errors.DB{Err: err}
	}

	return res, nil
}

func (s Services) Delete(ctx *gofr.Context, id string) error {
	if !idCheck(id) {
		return errors.Error("validation error")
	}

	err := s.Store.Delete(ctx, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func check(c models.Cat) bool {
	if !idCheck(c.ID) {
		return false
	}

	if c.Name == "" || c.Age <= 0 {
		return false
	}

	return true
}

func idCheck(i string) bool {
	id, err := strconv.Atoi(i)

	if err != nil {
		return false
	}

	if id <= 0 {
		return false
	}

	return true
}
