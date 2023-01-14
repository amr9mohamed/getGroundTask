package guests_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	guestsDef "github.com/getground/tech-tasks/backend/definitions/guests"
	tablesDef "github.com/getground/tech-tasks/backend/definitions/tables"
	"github.com/getground/tech-tasks/backend/pkg/database"
	"github.com/getground/tech-tasks/backend/pkg/modules/guests"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"regexp"
	"testing"
	"time"
)

type repoMocks struct {
	db      *sql.DB
	sqlMock sqlmock.Sqlmock
}

func setupIntegrationRepo(t *testing.T) (guestsDef.Repository, repoMocks) {
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
	r := guests.NewRepository(gDB)
	return r, repoMocks{
		db:      db,
		sqlMock: m,
	}
}

func TestRepository_Create(t *testing.T) {
	t.Run(
		"error create guest", func(t *testing.T) {

		},
	)

	t.Run(
		"error update table", func(t *testing.T) {

		},
	)

	t.Run(
		"success", func(t *testing.T) {

		},
	)
}

func TestRepository_GetByName(t *testing.T) {
	t.Run(
		"error", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			name := "test"

			//	mocks
			q := "SELECT * FROM `guests` WHERE name = ? AND time_arrived IS NULL ORDER BY `guests`.`name` LIMIT 1"
			m.sqlMock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(name).WillReturnError(errors.New("record not found"))

			//	method call
			res, err := repo.GetByName(name)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			// test data
			name := "test"
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 10,
				TimeArrived:  nil,
			}

			//	mocks
			q := "SELECT * FROM `guests` WHERE name = ? AND time_arrived IS NULL ORDER BY `guests`.`name` LIMIT 1"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(name).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"name", "table_id", "accompanying",
							"time_arrived",
						},
					).AddRow(g.Name, g.TableID, g.Accompanying, g.TimeArrived),
				)

			//	method call
			res, err := repo.GetByName(name)

			//	assert
			assert.NoError(t, err)
			assert.Equal(t, g, res)
		},
	)
}

func TestRepository_GetGuestList(t *testing.T) {
	t.Run(
		"error arrived", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			//	mocks
			q := "SELECT * FROM `guests` WHERE time_arrived IS NOT NULL"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs().
				WillReturnError(errors.New("record not found"))

			//	method call
			res, err := repo.GetGuestList(true)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"success arrived", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

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

			//	mocks
			q := "SELECT * FROM `guests` WHERE time_arrived IS NOT NULL"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs().
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"name", "table_id", "accompanying",
							"time_arrived",
						},
					).AddRow(gs[0].Name, gs[0].TableID, gs[0].Accompanying, gs[0].TimeArrived),
				)

			//	method call
			res, err := repo.GetGuestList(true)

			//	assert
			assert.NoError(t, err)
			assert.NotEmpty(t, res)
		},
	)

	t.Run(
		"error not arrived", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

			//	mocks
			q := "SELECT * FROM `guests`"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs().
				WillReturnError(errors.New("record not found"))

			//	method call
			res, err := repo.GetGuestList(false)

			//	assert
			assert.Error(t, err)
			assert.Empty(t, res)
		},
	)

	t.Run(
		"success not arrived", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)
			defer m.db.Close()

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

			//	mocks
			q := "SELECT * FROM `guests`"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs().
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"name", "table_id", "accompanying",
							"time_arrived",
						},
					).AddRow(gs[0].Name, gs[0].TableID, gs[0].Accompanying, gs[0].TimeArrived),
				)

			//	method call
			res, err := repo.GetGuestList(false)

			//	assert
			assert.NoError(t, err)
			assert.NotEmpty(t, res)
		},
	)
}

func TestRepository_CheckIn(t *testing.T) {
}

func TestRepository_CheckOut(t *testing.T) {
	t.Run(
		"guest not found", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)

			// test data
			name := "test"

			//	mocks
			q := "SELECT * FROM `guests` WHERE name = ? AND time_arrived IS NOT NULL ORDER BY `guests`.`name` LIMIT 1"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(name).
				WillReturnError(errors.New("guest not found"))

			//	method call
			err := repo.CheckOut(name)

			//	assert
			assert.Error(t, err)
		},
	)

	t.Run(
		"table not found", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)

			// test data
			name := "test"
			timeArrived := time.Now()
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 10,
				TimeArrived:  &timeArrived,
			}

			//	mocks
			guestQuery := "SELECT * FROM `guests` WHERE name = ? AND time_arrived IS NOT NULL ORDER BY `guests`.`name` LIMIT 1"
			tableQuery := "SELECT * FROM `tables` WHERE `tables`.`id` = ? ORDER BY `tables`.`id` LIMIT 1"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(guestQuery)).
				WithArgs(name).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"name", "table_id", "accompanying",
							"time_arrived",
						},
					).AddRow(g.Name, g.TableID, g.Accompanying, g.TimeArrived),
				)
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(tableQuery)).
				WithArgs(g.TableID).
				WillReturnError(errors.New("table not found"))

			//	method call
			err := repo.CheckOut(name)

			//	assert
			assert.Error(t, err)
		},
	)

	t.Run(
		"update table fail", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)

			// test data
			name := "test"
			timeArrived := time.Now()
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 10,
				TimeArrived:  &timeArrived,
			}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   10,
				EmptySeats: 10,
			}

			//	mocks
			guestQuery := "SELECT * FROM `guests` WHERE name = ? AND time_arrived IS NOT NULL ORDER BY `guests`.`name` LIMIT 1"
			tableQuery := "SELECT * FROM `tables` WHERE `tables`.`id` = ? ORDER BY `tables`.`id` LIMIT 1"
			updateTable := " UPDATE `tables` SET `empty_seats`=? WHERE `tables`.`id` = ?"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(guestQuery)).
				WithArgs(name).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"name", "table_id", "accompanying",
							"time_arrived",
						},
					).AddRow(g.Name, g.TableID, g.Accompanying, g.TimeArrived),
				)
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(tableQuery)).
				WithArgs(g.TableID).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "capacity", "empty_seats"}).AddRow(
						tbl.ID,
						tbl.Capacity, tbl.EmptySeats,
					),
				)
			m.sqlMock.ExpectBegin()
			m.sqlMock.ExpectExec(regexp.QuoteMeta(updateTable)).WithArgs(
				tbl.EmptySeats+g.Accompanying+1,
				tbl.ID,
			).WillReturnError(errors.New("error updating table"))
			m.sqlMock.ExpectRollback()

			//	method call
			err := repo.CheckOut(name)

			//	assert
			assert.Error(t, err)
		},
	)

	t.Run(
		"success", func(t *testing.T) {
			// setup
			repo, m := setupIntegrationRepo(t)

			// test data
			name := "test"
			timeArrived := time.Now()
			g := guestsDef.Guest{
				Name:         "test",
				TableID:      1,
				Accompanying: 10,
				TimeArrived:  &timeArrived,
			}
			tbl := tablesDef.Table{
				ID:         1,
				Capacity:   10,
				EmptySeats: 10,
			}

			//	mocks
			guestQuery := "SELECT * FROM `guests` WHERE name = ? AND time_arrived IS NOT NULL ORDER BY `guests`.`name` LIMIT 1"
			tableQuery := "SELECT * FROM `tables` WHERE `tables`.`id` = ? ORDER BY `tables`.`id` LIMIT 1"
			updateTable := " UPDATE `tables` SET `empty_seats`=? WHERE `tables`.`id` = ?"
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(guestQuery)).
				WithArgs(name).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"name", "table_id", "accompanying",
							"time_arrived",
						},
					).AddRow(g.Name, g.TableID, g.Accompanying, g.TimeArrived),
				)
			m.sqlMock.
				ExpectQuery(regexp.QuoteMeta(tableQuery)).
				WithArgs(g.TableID).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "capacity", "empty_seats"}).AddRow(
						tbl.ID,
						tbl.Capacity, tbl.EmptySeats,
					),
				)
			m.sqlMock.ExpectBegin()
			m.sqlMock.ExpectExec(regexp.QuoteMeta(updateTable)).WithArgs(
				tbl.EmptySeats+g.Accompanying+1,
				tbl.ID,
			).WillReturnResult(sqlmock.NewResult(1, 1))
			m.sqlMock.ExpectCommit()

			//	method call
			err := repo.CheckOut(name)

			//	assert
			assert.NoError(t, err)
		},
	)
}
