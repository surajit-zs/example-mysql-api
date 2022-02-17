package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"mysql-api/models"
)

type Store interface {
	Get(ctx *gofr.Context) ([]models.Cat, error)
	Create(ctx *gofr.Context, c models.Cat) (models.Cat, error)
}
