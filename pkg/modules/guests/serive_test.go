package guests_test

import (
	"errors"
	guestsDef "github.com/getground/tech-tasks/backend/definitions/guests"
	tablesDef "github.com/getground/tech-tasks/backend/definitions/tables"
	guestsMocks "github.com/getground/tech-tasks/backend/mocks/definitions/guests"
	tableMocks "github.com/getground/tech-tasks/backend/mocks/definitions/tables"
	"github.com/getground/tech-tasks/backend/pkg/modules/guests"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type serviceMocks struct {
	repo         *guestsMocks.Repository
	tableService *tableMocks.Service
}

func setupService() (guests.Service, serviceMocks) {
	repo := new(guestsMocks.Repository)
	tblService := new(tableMocks.Service)
	service := guests.NewService(repo, tblService)
	mocks := serviceMocks{repo, tblService}
	return service, mocks
}

func TestService_Create(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"table not found", func(t *testing.T) {
			// test data
			req := guestsDef.CreateRequest{Name: "test", Table: 1, Accompanying: 1}

			//	mocks
			m.tableService.On("GetByID", req.Table).Return(tablesDef.Table{}, errors.New("table not found")).Once()

			//	method call
			res, err := service.Create(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.tableService.AssertExpectations(t)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"no capacity", func(t *testing.T) {
			// test data
			req := guestsDef.CreateRequest{Name: "test", Table: 1, Accompanying: 10}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   5,
				EmptySeats: 5,
			}

			//	mocks
			m.tableService.On("GetByID", req.Table).Return(tbl, nil).Once()

			//	method call
			res, err := service.Create(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.tableService.AssertExpectations(t)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"repo error", func(t *testing.T) {
			// test data
			req := guestsDef.CreateRequest{Name: "test", Table: 1, Accompanying: 1}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   5,
				EmptySeats: 5,
			}

			//	mocks
			m.tableService.On("GetByID", req.Table).Return(tbl, nil).Once()
			m.repo.On("Create", req, tbl.Capacity-req.Accompanying-1).Return(
				errors.New(
					"error adding guest to guest list",
				),
			).Once()

			//	method call
			res, err := service.Create(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.tableService.AssertExpectations(t)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// test data
			req := guestsDef.CreateRequest{Name: "test", Table: 1, Accompanying: 1}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   5,
				EmptySeats: 5,
			}

			//	mocks
			m.tableService.On("GetByID", req.Table).Return(tbl, nil).Once()
			m.repo.On("Create", req, tbl.Capacity-req.Accompanying-1).Return(nil).Once()

			//	method call
			res, err := service.Create(req)

			//	assert
			assert.NoError(t, err)
			assert.NotEmpty(t, res)
			m.tableService.AssertExpectations(t)
			m.repo.AssertExpectations(t)
		},
	)
}

func TestService_GetGuestList(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"repo error", func(t *testing.T) {
			//	mocks
			m.repo.On("GetGuestList", false).Return([]guestsDef.Guest{}, errors.New("error retrieving")).Once()

			//	method call
			res, err := service.GetGuestList()

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// test data
			gs := []guestsDef.Guest{
				{
					Name:         "test",
					TableID:      1,
					Accompanying: 10,
				},
			}
			listDto := guestsDef.ListDTO{
				Guests: []guestsDef.GuestListDTO{
					{
						Name:         "test",
						Table:        1,
						Accompanying: 10,
					},
				},
			}
			//	mocks
			m.repo.On("GetGuestList", false).Return(gs, nil).Once()

			//	method call
			res, err := service.GetGuestList()

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, listDto, res)
			m.repo.AssertExpectations(t)
		},
	)
}

func TestService_GetGuests(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"repo error", func(t *testing.T) {
			//	mocks
			m.repo.On("GetGuestList", true).Return([]guestsDef.Guest{}, errors.New("error retrieving")).Once()

			//	method call
			res, err := service.GetGuests()

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// test data
			timeArrived := time.Now()
			gs := []guestsDef.Guest{
				{
					Name:         "test",
					TableID:      1,
					Accompanying: 10,
					TimeArrived:  &timeArrived,
				},
			}

			dto := guestsDef.DTO{
				Guests: []guestsDef.GuestDTO{
					{
						Name:         "test",
						Accompanying: 10,
						TimeArrived:  timeArrived.String(),
					},
				},
			}
			//	mocks
			m.repo.On("GetGuestList", true).Return(gs, nil).Once()

			//	method call
			res, err := service.GetGuests()

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, dto, res)
			m.repo.AssertExpectations(t)
		},
	)
}

func TestService_CheckIn(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"guest not found", func(t *testing.T) {
			//	test data
			req := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 10,
			}

			//	mocks
			m.repo.On("GetByName", req.Name).Return(guestsDef.Guest{}, errors.New("guest not invited")).Once()

			//	method call
			res, err := service.CheckIn(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"table not found", func(t *testing.T) {
			//	test data
			req := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 10,
			}
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 10,
				TimeArrived:  nil,
			}
			tbl := tablesDef.Table{}

			//	mocks
			m.repo.On("GetByName", req.Name).Return(g, nil).Once()
			m.tableService.On("GetByID", g.TableID).Return(tbl, errors.New("table not found")).Once()

			//	method call
			res, err := service.CheckIn(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"extra accompanying than expected", func(t *testing.T) {
			//	test data
			req := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 10,
			}
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 4,
				TimeArrived:  nil,
			}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   5,
				EmptySeats: 5,
			}

			//	mocks
			m.repo.On("GetByName", req.Name).Return(g, nil).Once()
			m.tableService.On("GetByID", g.TableID).Return(tbl, nil).Once()

			//	method call
			res, err := service.CheckIn(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"repo error", func(t *testing.T) {
			//	test data
			req := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 4,
			}
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 4,
				TimeArrived:  nil,
			}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   5,
				EmptySeats: 5,
			}

			//	mocks
			m.repo.On("GetByName", req.Name).Return(g, nil).Once()
			m.tableService.On("GetByID", g.TableID).Return(tbl, nil).Once()
			m.repo.On("CheckIn", req, g, tbl).Return(errors.New("error checking user in")).Once()

			//	method call
			res, err := service.CheckIn(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			//	test data
			req := guestsDef.CheckInRequest{
				Name:         "test",
				Accompanying: 4,
			}
			checkInRes := guestsDef.CheckInResponse{
				Name: "test",
			}
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 4,
				TimeArrived:  nil,
			}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   5,
				EmptySeats: 5,
			}

			//	mocks
			m.repo.On("GetByName", req.Name).Return(g, nil).Once()
			m.tableService.On("GetByID", g.TableID).Return(tbl, nil).Once()
			m.repo.On("CheckIn", req, g, tbl).Return(nil).Once()

			//	method call
			res, err := service.CheckIn(req)

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, checkInRes, res)
		},
	)
}

func TestService_CheckOut(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"success", func(t *testing.T) {
			// test data
			name := "test"

			//	mocks
			m.repo.On("CheckOut", name).Return(nil).Once()

			//	method call
			err := service.CheckOut(name)

			//	assert
			assert.NoError(t, err)
			m.repo.AssertExpectations(t)
		},
	)
}
