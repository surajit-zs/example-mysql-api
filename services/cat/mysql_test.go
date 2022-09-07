package cat

import (
	"context"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"

	"mysql-api/models"
	"mysql-api/store"
)

//nolint:dupl //they are defined separately for different test
func TestServices_Create(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}
	checkCat := models.Cat{
		ID:   "0",
		Name: "",
		Age:  0,
	}
	checkCat2 := models.Cat{
		ID:   "10",
		Name: "",
		Age:  0,
	}

	tests := []struct {
		desc string
		req  models.Cat
		mock []*gomock.Call
		res  models.Cat
		err  error
	}{{desc: "success", req: catModel,
		mock: []*gomock.Call{mockStore.EXPECT().Create(gomock.Any(), gomock.Any()).Return(catModel, nil)},
		res:  catModel, err: nil},
		{desc: "id is 0", req: checkCat, res: models.Cat{}, err: errors.Error("validation error")},
		{desc: "name is empty", req: checkCat2, res: models.Cat{}, err: errors.Error("validation error")},
		{desc: "error", req: catModel,
			mock: []*gomock.Call{mockStore.EXPECT().Create(gomock.Any(), gomock.Any()).Return(models.Cat{}, errors.Error("DB error"))},
			res:  models.Cat{}, err: errors.DB{Err: errors.Error("DB error")}},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.Create(ctx, tc.req)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %s,Got : %s ", i+1, tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.res) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.res, res)
			}
		})
	}
}

func TestServices_Get(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}
	tests := []struct {
		desc string
		mock []*gomock.Call
		res  []models.Cat
		err  error
	}{{desc: "success",
		mock: []*gomock.Call{mockStore.EXPECT().Get(gomock.Any()).Return([]models.Cat{catModel}, nil)},
		res:  []models.Cat{catModel}, err: nil},
		{desc: "failed",
			mock: []*gomock.Call{mockStore.EXPECT().Get(gomock.Any()).Return(nil, errors.Error("db error"))},
			res:  nil, err: errors.DB{Err: errors.Error("db error")}},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.Get(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %s,Got : %s ", i+1, tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.res) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.res, res)
			}
		})
	}
}

func Test_check(t *testing.T) {
	checkCat := models.Cat{
		ID:   "0",
		Name: "",
		Age:  0,
	}
	checkCat2 := models.Cat{
		ID:   "10",
		Name: "",
		Age:  0,
	}
	checkCat3 := models.Cat{
		ID:   "10",
		Name: "test_name",
		Age:  2,
	}
	checkCat5 := models.Cat{
		ID:   "hjg",
		Name: "test_name",
		Age:  2,
	}

	tests := []struct {
		desc   string
		input  models.Cat
		output bool
	}{{"success", checkCat3, true},
		{"empty name", checkCat2, false},
		{"id is negative", checkCat, false},
		{"id is not int", checkCat5, false},
	}

	for i, tc := range tests {
		res := check(tc.input)
		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.output, res)
		}
	}
}

func TestServices_Delete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}
	checkCat := models.Cat{
		ID:   "0",
		Name: "",
		Age:  0,
	}

	tests := []struct {
		desc string
		id   string
		mock []*gomock.Call
		err  error
	}{{desc: "success", id: catModel.ID,
		mock: []*gomock.Call{mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)}, err: nil},
		{desc: "id is 0", id: checkCat.ID, err: errors.Error("validation error")},
		{desc: "db error", id: catModel.ID, mock: []*gomock.Call{mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).
			Return(errors.Error("db error"))}, err: errors.DB{Err: errors.Error("db error")}},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			err := mock.Delete(ctx, tc.id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %s,Got : %s ", i+1, tc.err, err)
			}
		})
	}
}

//nolint:dupl //they are defined separately for different test
func TestServices_Update(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}
	checkCat := models.Cat{
		ID:   "0",
		Name: "",
		Age:  0,
	}
	checkCat2 := models.Cat{
		ID:   "10",
		Name: "",
		Age:  0,
	}

	tests := []struct {
		desc string
		req  models.Cat
		mock []*gomock.Call
		res  models.Cat
		err  error
	}{{desc: "success", req: catModel,
		mock: []*gomock.Call{mockStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(catModel, nil)},
		res:  catModel, err: nil},
		{desc: "id is 0", req: checkCat, res: models.Cat{}, err: errors.Error("validation error")},
		{desc: "name is empty", req: checkCat2, res: models.Cat{}, err: errors.Error("validation error")},
		{desc: "error", req: catModel,
			mock: []*gomock.Call{mockStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(models.Cat{}, errors.Error("DB error"))},
			res:  models.Cat{}, err: errors.DB{Err: errors.Error("DB error")}},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.Update(ctx, tc.req)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %s,Got : %s ", i+1, tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.res) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.res, res)
			}
		})
	}
}

func TestServices_GetByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}
	checkCat := models.Cat{
		ID:   "0",
		Name: "",
		Age:  0,
	}
	checkCat2 := models.Cat{
		ID:   "10",
		Name: "",
		Age:  0,
	}

	tests := []struct {
		desc string
		id   string
		mock []*gomock.Call
		res  models.Cat
		err  error
	}{{desc: "success", id: catModel.ID,
		mock: []*gomock.Call{mockStore.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(catModel, nil)},
		res:  catModel, err: nil},

		{desc: "id is 0", id: checkCat.ID, res: models.Cat{}, err: errors.Error("validation error")},

		{desc: "error", id: checkCat2.ID,
			mock: []*gomock.Call{mockStore.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Cat{}, errors.Error("DB error"))},
			res:  models.Cat{}, err: errors.DB{Err: errors.Error("DB error")}},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetByID(ctx, tc.id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %s,Got : %s ", i+1, tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.res) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.res, res)
			}
		})
	}
}
