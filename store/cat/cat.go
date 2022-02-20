package cat

import (
	"database/sql"

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

func (c catStore) GetByID(ctx *gofr.Context, id string) (models.Cat, error) {
	var resp models.Cat

	query := "select * from cat where id=$1"
	err := ctx.DB().QueryRowContext(ctx, query, id).Scan(&resp.ID, &resp.Name, &resp.Age)

	if err == sql.ErrNoRows {
		return models.Cat{}, errors.EntityNotFound{Entity: "cat", ID: id}
	}

	if err != nil {
		return models.Cat{}, errors.DB{Err: err}
	}

	return resp, nil
}

func (c catStore) Create(ctx *gofr.Context, cat models.Cat) (models.Cat, error) {
	query := "Insert into cat values(?,?,?)"
	_, err := ctx.DB().ExecContext(ctx, query, cat.ID, cat.Name, cat.Age)

	if err != nil {
		return cat, errors.DB{Err: err}
	}

	return cat, nil
}

func (c catStore) Update(ctx *gofr.Context, cat models.Cat) (models.Cat, error) {
	query := "UPDATE cat SET name=$1, age=$2 WHERE id=$3"
	_, err := ctx.DB().ExecContext(ctx, query, cat.Name, cat.Age, cat.ID)

	if err == sql.ErrNoRows {
		return models.Cat{}, errors.EntityNotFound{Entity: "cat", ID: cat.ID}
	}

	if err != nil {
		return models.Cat{}, errors.DB{Err: err}
	}

	return cat, nil
}

func (c catStore) Delete(ctx *gofr.Context, id string) error {
	query := "DELETE FROM cat where id=$1"
	_, err := ctx.DB().ExecContext(ctx, query, id)

	if err == sql.ErrNoRows {
		return errors.EntityNotFound{Entity: "cat", ID: id}
	}

	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
