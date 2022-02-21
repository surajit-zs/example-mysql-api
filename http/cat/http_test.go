package cat

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"

	"mysql-api/models"
	"mysql-api/services"
)

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := services.NewMockServices(ctrl)
	h := Handler{s}
	app := gofr.New()

	test1 := `{"id":"1","name":"test_name","age":2}`
	test2 := `{"id":"1","name":"test_name","age":2,"TEST";"TEST"}`
	test3 := `{"id":"0","name":"test_name","age":2}`
	test4 := `{"id":"4","name":"test_name","age":2}`

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}

	tests := []struct {
		desc string
		req  []byte
		mock []*gomock.Call
		err  error
	}{{desc: "success", req: []byte(test1),
		mock: []*gomock.Call{s.EXPECT().Create(gomock.Any(), gomock.Any()).Return(catModel, nil)}, err: nil},
		{desc: "ID is 0", req: []byte(test3), err: errors.InvalidParam{Param: []string{"id"}}},
		{desc: "unmarsel", req: []byte(test2), err: errors.InvalidParam{Param: []string{"body"}}},
		{desc: "db error", req: []byte(test4),
			mock: []*gomock.Call{s.EXPECT().Create(gomock.Any(), gomock.Any()).Return(models.Cat{},
				errors.Error("Internal DB error"))}, err: errors.Error("Internal DB error")}}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "https://cat", bytes.NewReader(tc.req))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			_, err := h.Create(ctx)

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.err, err)
			}
		})
	}
}

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := services.NewMockServices(ctrl)
	h := Handler{s}
	app := gofr.New()

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  2,
	}

	testRes := models.Response{
		Cat:        []models.Cat{catModel},
		Massage:    "Success",
		StatusCode: http.StatusOK,
	}

	tests := []struct {
		desc string
		res  interface{}
		mock []*gomock.Call
		err  error
	}{{desc: "success", res: testRes,
		mock: []*gomock.Call{s.EXPECT().Get(gomock.Any()).Return([]models.Cat{catModel}, nil)}},
		{desc: "error", res: nil,
			mock: []*gomock.Call{s.EXPECT().Get(gomock.Any()).Return([]models.Cat{}, errors.Error("Internal DB error"))},
			err:  errors.DB{Err: errors.Error("Internal DB error")}},
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "https://cat", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			result, err := h.Get(ctx)

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.err, err)
			}

			if !reflect.DeepEqual(tc.res, result) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.res, result)
			}
		})
	}
}

func TestHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := services.NewMockServices(ctrl)
	h := Handler{s}
	app := gofr.New()

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  2,
	}

	testRes := models.Response{
		Cat:        catModel,
		Massage:    "successful",
		StatusCode: http.StatusOK,
	}

	tests := []struct {
		desc string
		id   string
		res  interface{}
		mock []*gomock.Call
		err  error
	}{{desc: "success", id: catModel.ID, res: testRes,
		mock: []*gomock.Call{s.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(catModel, nil)}},

		{desc: "db error", id: catModel.ID, res: nil,
			mock: []*gomock.Call{s.EXPECT().GetByID(gomock.Any(), catModel.ID).Return(models.Cat{}, errors.Error("Internal DB error"))},
			err: errors.EntityNotFound{
				Entity: "cat",
				ID:     catModel.ID,
			}},

		{desc: "id is not int", id: "fsef", err: errors.InvalidParam{
			Param: []string{"id"},
		}},

		{desc: "id is empty", id: "", err: errors.MissingParam{Param: []string{"id"}}},
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "https://cat", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})
			result, err := h.GetByID(ctx)

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.err, err)
			}

			if !reflect.DeepEqual(tc.res, result) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.res, result)
			}
		})
	}
}

func TestHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := services.NewMockServices(ctrl)
	h := Handler{s}
	app := gofr.New()

	test1 := `{"id":"1","name":"test_name","age":2}`
	test2 := `{"id":"1","name":"test_name","age":2,"TEST";"TEST"}`
	test3 := `{"id":"0","name":"test_name","age":2}`
	test4 := `{"id":"4","name":"test_name","age":2}`

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  1,
	}

	tests := []struct {
		desc string
		id   string
		req  []byte
		mock []*gomock.Call
		err  error
	}{{desc: "success", id: "1", req: []byte(test1),
		mock: []*gomock.Call{s.EXPECT().Update(gomock.Any(), gomock.Any()).Return(catModel, nil)}, err: nil},

		{desc: "id is empty", id: "", err: errors.MissingParam{Param: []string{"id"}}},

		{desc: "id is not int", id: "sef", err: errors.InvalidParam{
			Param: []string{"id"},
		}},

		{desc: "ID is 0", id: "0", req: []byte(test3), err: errors.InvalidParam{Param: []string{"id"}}},

		{desc: "unmarsel", id: "1", req: []byte(test2), err: errors.InvalidParam{Param: []string{"body"}}},

		{desc: "db error", id: "1", req: []byte(test4),
			mock: []*gomock.Call{s.EXPECT().Update(gomock.Any(), gomock.Any()).Return(models.Cat{},
				errors.Error("Internal DB error"))}, err: errors.Error("Internal DB error")}}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "https://cat", bytes.NewReader(tc.req))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})
			_, err := h.Update(ctx)

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.err, err)
			}
		})
	}
}

func TestHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := services.NewMockServices(ctrl)
	h := Handler{s}
	app := gofr.New()

	catModel := models.Cat{
		ID:   "1",
		Name: "test_name",
		Age:  2,
	}

	testRes := models.Response{
		Cat:        nil,
		Massage:    "successful",
		StatusCode: http.StatusOK,
	}

	tests := []struct {
		desc string
		id   string
		res  interface{}
		mock []*gomock.Call
		err  error
	}{{desc: "success", id: catModel.ID, res: testRes,
		mock: []*gomock.Call{s.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)}},

		{desc: "db error", id: catModel.ID, res: nil,
			mock: []*gomock.Call{s.EXPECT().Delete(gomock.Any(), catModel.ID).Return(errors.Error("Internal DB error"))},
			err:  errors.DB{Err: errors.Error("Internal DB error")}},

		{desc: "id is not int", id: "fsef", err: errors.InvalidParam{
			Param: []string{"id"},
		}},

		{desc: "id is empty", id: "", err: errors.MissingParam{Param: []string{"id"}}},
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "https://cat", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})
			_, err := h.Delete(ctx)

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("test :%v Expected : %v,Got : %v ", i+1, tc.err, err)
			}
		})
	}
}
