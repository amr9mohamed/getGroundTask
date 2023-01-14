package tables_test

import (
	"errors"
	tablesDef "github.com/getground/tech-tasks/backend/definitions/tables"
	tableMocks "github.com/getground/tech-tasks/backend/mocks/definitions/tables"
	"github.com/getground/tech-tasks/backend/pkg/modules/tables"
	"github.com/stretchr/testify/assert"
	"testing"
)

type serviceMocks struct {
	repo *tableMocks.Repository
}

func setupService() (tables.Service, serviceMocks) {
	repo := new(tableMocks.Repository)
	service := tables.NewService(repo)
	mocks := serviceMocks{repo}
	return service, mocks
}

func TestService_Create(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"repository error", func(t *testing.T) {
			//	test date
			createReq := tablesDef.CreateRequest{Capacity: 10}
			//	mocks
			m.repo.On("Create", createReq).Return(tablesDef.Table{}, errors.New("error creating table")).Once()

			//	method call
			res, err := service.Create(createReq)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			//	test date
			createReq := tablesDef.CreateRequest{Capacity: 10}
			tbl := tablesDef.Table{ID: 1, Capacity: 10}
			createRes := tablesDef.CreateResponse{ID: 1, Capacity: 10}
			//	mocks
			m.repo.On("Create", createReq).Return(tbl, nil).Once()

			//	method call
			res, err := service.Create(createReq)

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, createRes, res)
			m.repo.AssertExpectations(t)
		},
	)
}

func TestService_GetByID(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"repo error", func(t *testing.T) {
			// test data
			id := uint(1)
			//	mocks
			m.repo.On("GetByID", id).Return(tablesDef.Table{}, errors.New("table not found")).Once()

			//	method call
			res, err := service.GetByID(id)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
			m.repo.AssertExpectations(t)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// test data
			id := uint(1)
			tbl := tablesDef.Table{ID: 1, Capacity: 10}
			//	mocks
			m.repo.On("GetByID", id).Return(tbl, nil).Once()

			//	method call
			res, err := service.GetByID(id)

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, tbl, res)
			m.repo.AssertExpectations(t)
		},
	)
}

func TestService_CountEmptySeats(t *testing.T) {
	// setup
	service, m := setupService()
	t.Run(
		"success", func(t *testing.T) {
			//	mocks
			m.repo.On("CountEmptySeats").Return(10).Once()

			//	method call
			count := service.CountEmptySeats()

			//	assert
			assert.Equal(t, count, 10)
			m.repo.AssertExpectations(t)
		},
	)
}
