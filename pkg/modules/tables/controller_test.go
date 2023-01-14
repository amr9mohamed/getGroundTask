package tables_test

import (
	"encoding/json"
	"errors"
	tableDef "github.com/getground/tech-tasks/backend/definitions/tables"
	tableMocks "github.com/getground/tech-tasks/backend/mocks/definitions/tables"
	"github.com/getground/tech-tasks/backend/pkg/modules/tables"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ctrlMocks struct {
	handler tables.Handler
	service *tableMocks.Service
}

func setupController() (*gin.Engine, tables.Controller, ctrlMocks) {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	handler := tables.NewHandler()
	service := new(tableMocks.Service)
	ctrl := tables.NewController(handler, service)
	mocks := ctrlMocks{handler, service}

	return r, ctrl, mocks
}

func TestController_Create(t *testing.T) {
	//	setup
	r, ctrl, m := setupController()
	r.POST("/tables", ctrl.Create)
	t.Run(
		"error in handler", func(t *testing.T) {
			//	test data
			createReq := tableDef.CreateRequest{
				Capacity: 0,
			}

			//	request
			body, err := json.Marshal(&createReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/tables", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// assert
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"error in service", func(t *testing.T) {
			//	test data
			createReq := tableDef.CreateRequest{
				Capacity: 10,
			}

			// mocks
			m.service.On("Create", createReq).Return(
				tableDef.CreateResponse{}, errors.New("internal error creating table"),
			).Once()

			//	request
			body, err := json.Marshal(&createReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}
			req, err := http.NewRequest(http.MethodPost, "/tables", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// assert
			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			//	test data
			createReq := tableDef.CreateRequest{
				Capacity: 10,
			}

			createRes := tableDef.CreateResponse{
				ID:       1,
				Capacity: 10,
			}

			// mocks
			m.service.On("Create", createReq).Return(createRes, nil).Once()

			//	request
			body, err := json.Marshal(&createReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}
			req, err := http.NewRequest(http.MethodPost, "/tables", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// expectation
			expected := `{"id":1,"capacity":10}`

			// assert
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, expected, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}

func TestController_CountEmptySeats(t *testing.T) {
	//	setup
	r, ctrl, m := setupController()
	t.Run(
		"success", func(t *testing.T) {
			//	mocks
			m.service.On("CountEmptySeats").Return(10)

			//	request
			r.GET("/empty_seats", ctrl.CountEmptySeats)
			req, err := http.NewRequest(http.MethodGet, "/empty_seats", http.NoBody)
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			expected := `{"seats_empty":10}`

			// assert
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, expected, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}
