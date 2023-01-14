package tables_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	tablesDef "github.com/getground/tech-tasks/backend/definitions/tables"
	"github.com/getground/tech-tasks/backend/pkg/database"
	"github.com/getground/tech-tasks/backend/pkg/modules/tables"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"regexp"
	"testing"
)

type repoMocks struct {
	db      *sql.DB
	sqlMock sqlmock.Sqlmock
}

func setupIntegrationRepo(t *testing.T) (tablesDef.Repository, repoMocks) {
	db, m, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	msc := mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true})
	gDB, err := database.NewDatabaseForTests(msc)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating grom database connection", err)
	}
	m.MatchExpectationsInOrder(false)
	r := tables.NewRepository(gDB)
	return r, repoMocks{
		db:      db,
		sqlMock: m,
	}
}

func TestRepository_Create(t *testing.T) {
	t.Run(
		"error", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			req := tablesDef.CreateRequest{Capacity: 10}
			// mocks
			q := "INSERT INTO `tables` (`capacity`,`empty_seats`) VALUES (?,?)"
			m.sqlMock.
				ExpectExec(regexp.QuoteMeta(q)).
				WithArgs(req.Capacity, req.Capacity).
				WillReturnError(errors.New("table not found"))

			// method call
			tbl, err := repo.Create(req)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, tbl)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			req := tablesDef.CreateRequest{Capacity: 10}
			// mocks
			m.sqlMock.ExpectBegin()
			q := "INSERT INTO `tables` (`capacity`,`empty_seats`) VALUES (?,?)"
			m.sqlMock.
				ExpectExec(regexp.QuoteMeta(q)).
				WithArgs(int(req.Capacity), int(req.Capacity)).
				WillReturnResult(sqlmock.NewResult(1, 1))
			m.sqlMock.ExpectCommit()

			// method call
			res, err := repo.Create(req)

			// expectations
			expecteTable := tablesDef.Table{
				ID:         1,
				Capacity:   10,
				EmptySeats: 10,
			}

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, expecteTable, res)
		},
	)
}

func TestRepository_GetByID(t *testing.T) {
	t.Run(
		"error", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			id := uint(1)

			//	mocks
			q := "SELECT * FROM `tables` WHERE `tables`.`id` = ? ORDER BY `tables`.`id` LIMIT 1"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(id).
				WillReturnError(errors.New("table not found"))

			//	method call
			res, err := repo.GetByID(id)

			// expectation
			expectedTable := tablesDef.Table{}

			//	assert
			assert.Error(t, err)
			assert.Equal(t, expectedTable, res)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			id := uint(1)

			// columns
			tColumns := []string{"id", "capacity", "empty_seats"}

			// test data
			tValues := []driver.Value{1, 10, 10}
			//	mocks
			q := "SELECT * FROM `tables` WHERE `tables`.`id` = ? ORDER BY `tables`.`id` LIMIT 1"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows(tColumns).AddRow(tValues...))

			//	method call
			res, err := repo.GetByID(id)

			// expectation
			expectedTable := tablesDef.Table{
				ID:         1,
				Capacity:   10,
				EmptySeats: 10,
			}

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, expectedTable, res)
		},
	)
}

func TestRepository_CountEmptySeats(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			sum := int(10)
			//	mocks
			q := "SELECT SUM(empty_seats) FROM `tables`"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WillReturnRows(sqlmock.NewRows([]string{"sum(empty_seats)"}).AddRow(sum))

			//	method call
			res := repo.CountEmptySeats()

			//	assert
			assert.Equal(t, sum, res)
		},
	)
}
