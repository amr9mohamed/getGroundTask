// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	guests "github.com/getground/tech-tasks/backend/definitions/guests"
	mock "github.com/stretchr/testify/mock"

	tables "github.com/getground/tech-tasks/backend/definitions/tables"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckIn provides a mock function with given fields: request, guest, table
func (_m *Repository) CheckIn(request guests.CheckInRequest, guest guests.Guest, table tables.Table) error {
	ret := _m.Called(request, guest, table)

	var r0 error
	if rf, ok := ret.Get(0).(func(guests.CheckInRequest, guests.Guest, tables.Table) error); ok {
		r0 = rf(request, guest, table)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckOut provides a mock function with given fields: name
func (_m *Repository) CheckOut(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: request, tableCapacity
func (_m *Repository) Create(request guests.CreateRequest, tableCapacity int64) error {
	ret := _m.Called(request, tableCapacity)

	var r0 error
	if rf, ok := ret.Get(0).(func(guests.CreateRequest, int64) error); ok {
		r0 = rf(request, tableCapacity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByName provides a mock function with given fields: name
func (_m *Repository) GetByName(name string) (guests.Guest, error) {
	ret := _m.Called(name)

	var r0 guests.Guest
	if rf, ok := ret.Get(0).(func(string) guests.Guest); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(guests.Guest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGuestList provides a mock function with given fields: arrived
func (_m *Repository) GetGuestList(arrived bool) ([]guests.Guest, error) {
	ret := _m.Called(arrived)

	var r0 []guests.Guest
	if rf, ok := ret.Get(0).(func(bool) []guests.Guest); ok {
		r0 = rf(arrived)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]guests.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(arrived)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
