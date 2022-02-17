package cat

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"mysql-api/models"
	"mysql-api/store"
)

type catStore struct{}

func New() store.Store {
	return catStore{}
}

func (c catStore) Get(ctx *gofr.Context) ([]models.Cat, error) {
	query := "select * from cat"

	rows, err := ctx.DB().QueryContext(ctx, query)

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	var cats []models.Cat

	for rows.Next() {
		var c models.Cat

		err := rows.Scan(&c.ID, &c.Name, &c.Age)
		if err != nil {
			return nil, errors.DB{Err: errors.Error("scan error")}
		}

		cats = append(cats, c)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.DB{Err: err}
	}

	return cats, nil
}

func (c catStore) Create(ctx *gofr.Context, cat models.Cat) (models.Cat, error) {
	query := "Insert into cat values(?,?,?)"
	_, err := ctx.DB().ExecContext(ctx, query, cat.ID, cat.Name, cat.Age)

	if err != nil {
		return cat, errors.DB{Err: err}
	}

	return cat, nil
}
