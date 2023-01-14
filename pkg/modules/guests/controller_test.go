package guests_test

import (
	"encoding/json"
	"errors"
	"fmt"
	guestsDef "github.com/getground/tech-tasks/backend/definitions/guests"
	guestsMocks "github.com/getground/tech-tasks/backend/mocks/definitions/guests"
	"github.com/getground/tech-tasks/backend/pkg/modules/guests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ctrlMocks struct {
	handler guests.Handler
	service *guestsMocks.Service
}

func setupController() (*gin.Engine, guests.Controller, ctrlMocks) {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	handler := guests.NewHandler()
	service := new(guestsMocks.Service)
	ctrl := guests.NewController(handler, service)
	mocks := ctrlMocks{handler, service}

	return r, ctrl, mocks
}

func TestController_Create(t *testing.T) {
	t.Run(
		"handler error, name not sent", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.POST("/guest_list/:test", ctrl.Create)
			//	test data
			createRequest := guestsDef.CreateRequest{Name: ""}

			//	request
			body, err := json.Marshal(&createRequest)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/guest_list/:test", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			m.service.AssertExpectations(t)
		},
	)
	t.Run(
		"handler error", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.POST("/guest_list/:name", ctrl.Create)
			//	test data
			createRequest := guestsDef.CreateRequest{Name: ""}

			//	request
			body, err := json.Marshal(&createRequest)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/guest_list/:name", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"service error", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.POST("/guest_list/:name", ctrl.Create)
			//	test data
			createRequest := guestsDef.CreateRequest{
				Name:         "test",
				Table:        1,
				Accompanying: 10,
			}
			createResponse := guestsDef.CreateResponse{}

			// mocks
			m.service.On("Create", createRequest).Return(createResponse, errors.New("internal error")).Once()

			//	request
			body, err := json.Marshal(&createRequest)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/guest_list/:name", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.POST("/guest_list/:name", ctrl.Create)
			//	test data
			createRequest := guestsDef.CreateRequest{
				Name:         "test",
				Table:        1,
				Accompanying: 10,
			}
			createResponse := guestsDef.CreateResponse{
				Name: "test",
			}

			// mocks
			m.service.On("Create", createRequest).Return(createResponse, nil).Once()

			//	request
			body, err := json.Marshal(&createRequest)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/guest_list/:name", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// expectation
			expected := `{"name":"test"}`

			//	assert
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, expected, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}

func TestController_GetGuestList(t *testing.T) {
	//	setup
	r, ctrl, m := setupController()
	r.GET("/guest_list", ctrl.GetGuestList)
	t.Run(
		"service error", func(t *testing.T) {
			// mocks
			m.service.On("GetGuestList").Return(guestsDef.ListDTO{}, errors.New("internal error")).Once()

			//	request
			req, err := http.NewRequest(http.MethodGet, "/guest_list", http.NoBody)
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
			// test data
			res := guestsDef.ListDTO{
				Guests: []guestsDef.GuestListDTO{
					{
						Name:         "test",
						Table:        1,
						Accompanying: 10,
					},
				},
			}
			// mocks
			m.service.On("GetGuestList").Return(res, nil).Once()

			//	request
			req, err := http.NewRequest(http.MethodGet, "/guest_list", http.NoBody)
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// expectation
			expected := `{"guests":[{"name":"test","table":1,"accompanying_guests":10}]}`

			// assert
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, expected, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}

func TestController_GetGuests(t *testing.T) {
	//	setup
	r, ctrl, m := setupController()
	r.GET("/guests", ctrl.GetGuests)
	t.Run(
		"service error", func(t *testing.T) {
			// mocks
			m.service.On("GetGuests").Return(guestsDef.DTO{}, errors.New("internal error")).Once()

			//	request
			req, err := http.NewRequest(http.MethodGet, "/guests", http.NoBody)
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
			// test data
			res := guestsDef.DTO{
				Guests: []guestsDef.GuestDTO{
					{
						Name:         "test",
						Accompanying: 10,
						TimeArrived:  "14/1/223",
					},
				},
			}
			// mocks
			m.service.On("GetGuests").Return(res, nil).Once()

			//	request
			req, err := http.NewRequest(http.MethodGet, "/guests", http.NoBody)
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// expectation
			expected := `{"guests":[{"name":"test","accompanying_guests":10,"time_arrived":"14/1/223"}]}`

			// assert
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, expected, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}

func TestController_CheckIn(t *testing.T) {
	t.Run(
		"handler err, name not sent", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.PUT("/guests/:test", ctrl.CheckIn)

			//	test data
			checkInReq := guestsDef.CheckInRequest{Name: ""}

			//	request
			body, err := json.Marshal(&checkInReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPut, "/guests/:test", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"handler err", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.PUT("/guests/:name", ctrl.CheckIn)

			//	test data
			checkInReq := guestsDef.CheckInRequest{Name: ""}

			//	request
			body, err := json.Marshal(&checkInReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPut, "/guests/:name", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"service error", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.PUT("/guests/:name", ctrl.CheckIn)

			//	test data
			checkInReq := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 10,
			}
			checkInRes := guestsDef.CheckInResponse{}

			// mocks
			m.service.On("CheckIn", checkInReq).Return(checkInRes, errors.New("internal error")).Once()

			//	request
			body, err := json.Marshal(&checkInReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPut, "/guests/:name", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()
			r.PUT("/guests/:name", ctrl.CheckIn)

			//	test data
			checkInReq := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 10,
			}
			checkInRes := guestsDef.CheckInResponse{
				Name: "test",
			}

			// mocks
			m.service.On("CheckIn", checkInReq).Return(checkInRes, nil).Once()

			//	request
			body, err := json.Marshal(&checkInReq)
			if err != nil {
				t.Errorf("Error converting struct to json - test controller: %v\n", err)
			}

			req, err := http.NewRequest(http.MethodPut, "/guests/:name", strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			// expectation
			expected := `{"name":"test"}`

			//	assert
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, expected, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}

func TestController_CheckOut(t *testing.T) {
	t.Run(
		"handler err", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()

			//	request
			r.DELETE("/guests/:test", ctrl.CheckOut)
			req, err := http.NewRequest(http.MethodDelete, "/guests/:test", http.NoBody)
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"service error", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()

			//	test data
			name := "test"

			// mocks
			m.service.On("CheckOut", name).Return(errors.New("internal error")).Once()

			//	request
			r.DELETE("/guests/:name", ctrl.CheckOut)

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/guests/%s", name), http.NoBody)
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			m.service.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			//	setup
			r, ctrl, m := setupController()

			//	test data
			name := "test"

			// mocks
			m.service.On("CheckOut", name).Return(nil).Once()

			//	request
			r.DELETE("/guests/:name", ctrl.CheckOut)

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/guests/%s", name), nil)
			if err != nil {
				t.Errorf("Error requesting test controller: %v\n", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			//	assert
			assert.Equal(t, http.StatusNoContent, rr.Code)
			assert.Empty(t, rr.Body.String())
			m.service.AssertExpectations(t)
		},
	)
}
