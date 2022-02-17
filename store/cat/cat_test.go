package cat

import (
	"context"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"

	"mysql-api/models"
)

func TestCatStore_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	cat1 := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}

	query := "Insert into cat values(?,?,?)"
	tests := []struct {
		desc  string
		input models.Cat
		mock  []interface{}
		err   error
	}{
		{"success", cat1,
			[]interface{}{mock.ExpectExec(query).WithArgs(cat1.ID, cat1.Name, cat1.Age).
				WillReturnResult(sqlmock.NewResult(1, 1))}, nil},

		{"error", cat1, []interface{}{mock.ExpectExec(query).WithArgs(cat1.ID, cat1.Name, cat1.Age).
			WillReturnError(errors.Error("db error"))}, errors.DB{Err: errors.Error("db error")}},
	}

	g := gofr.Gofr{DataStore: datastore.DataStore{ORM: db}}
	ctx := gofr.NewContext(nil, nil, &g)
	ctx.Context = context.Background()
	store := New()

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := store.Create(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Test :%v Expected : %v,Got : %v ", i+1, tc.err, err)
			}
		})
	}
}

func TestCatStore_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	cat1 := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}

	query := "select * from cat"
	row := sqlmock.NewRows([]string{"id", "name", "age"}).AddRow("1", "test_name", 1)
	scanError := sqlmock.NewRows([]string{"id", "name", "age", "err"}).AddRow("1", 12, 1, "err")
	rowError := sqlmock.NewRows([]string{"id", "name", "err"}).AddRow("2", "test", 2).
		RowError(0, errors.Error("row error"))

	tests := []struct {
		desc   string
		mock   []interface{}
		output []models.Cat
		err    error
	}{
		{"success",
			[]interface{}{mock.ExpectQuery(query).WithArgs().WillReturnRows(row)},
			[]models.Cat{cat1}, nil},

		{"error", []interface{}{mock.ExpectQuery(query).WithArgs().
			WillReturnError(errors.Error("db error"))}, nil,
			errors.DB{Err: errors.Error("db error")}},

		{"scan error", []interface{}{mock.ExpectQuery(query).WithArgs().
			WillReturnRows(scanError)}, nil,
			errors.DB{Err: errors.Error("scan error")}},

		{"row error", []interface{}{mock.ExpectQuery(query).WithArgs().
			WillReturnRows(rowError)}, nil, errors.DB{Err: errors.Error("row error")}},
	}

	g := gofr.Gofr{DataStore: datastore.DataStore{ORM: db}}
	ctx := gofr.NewContext(nil, nil, &g)
	ctx.Context = context.Background()
	store := New()

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.Get(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Got : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.output) {
				t.Errorf("Expected : %v,Got : %v ", tc.output, res)
			}
		})
	}
}
