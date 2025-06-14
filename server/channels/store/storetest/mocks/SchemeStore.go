// Code generated by mockery v2.53.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// SchemeStore is an autogenerated mock type for the SchemeStore type
type SchemeStore struct {
	mock.Mock
}

// CountByScope provides a mock function with given fields: scope
func (_m *SchemeStore) CountByScope(scope string) (int64, error) {
	ret := _m.Called(scope)

	if len(ret) == 0 {
		panic("no return value specified for CountByScope")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(scope)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(scope)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(scope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountWithoutPermission provides a mock function with given fields: scope, permissionID, roleScope, roleType
func (_m *SchemeStore) CountWithoutPermission(scope string, permissionID string, roleScope model.RoleScope, roleType model.RoleType) (int64, error) {
	ret := _m.Called(scope, permissionID, roleScope, roleType)

	if len(ret) == 0 {
		panic("no return value specified for CountWithoutPermission")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.RoleScope, model.RoleType) (int64, error)); ok {
		return rf(scope, permissionID, roleScope, roleType)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.RoleScope, model.RoleType) int64); ok {
		r0 = rf(scope, permissionID, roleScope, roleType)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string, model.RoleScope, model.RoleType) error); ok {
		r1 = rf(scope, permissionID, roleScope, roleType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: schemeID
func (_m *SchemeStore) Delete(schemeID string) (*model.Scheme, error) {
	ret := _m.Called(schemeID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *model.Scheme
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Scheme, error)); ok {
		return rf(schemeID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Scheme); ok {
		r0 = rf(schemeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Scheme)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(schemeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: schemeID
func (_m *SchemeStore) Get(schemeID string) (*model.Scheme, error) {
	ret := _m.Called(schemeID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Scheme
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Scheme, error)); ok {
		return rf(schemeID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Scheme); ok {
		r0 = rf(schemeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Scheme)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(schemeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPage provides a mock function with given fields: scope, offset, limit
func (_m *SchemeStore) GetAllPage(scope string, offset int, limit int) ([]*model.Scheme, error) {
	ret := _m.Called(scope, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetAllPage")
	}

	var r0 []*model.Scheme
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*model.Scheme, error)); ok {
		return rf(scope, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.Scheme); ok {
		r0 = rf(scope, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Scheme)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(scope, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: schemeName
func (_m *SchemeStore) GetByName(schemeName string) (*model.Scheme, error) {
	ret := _m.Called(schemeName)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *model.Scheme
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Scheme, error)); ok {
		return rf(schemeName)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Scheme); ok {
		r0 = rf(schemeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Scheme)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(schemeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteAll provides a mock function with no fields
func (_m *SchemeStore) PermanentDeleteAll() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PermanentDeleteAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: scheme
func (_m *SchemeStore) Save(scheme *model.Scheme) (*model.Scheme, error) {
	ret := _m.Called(scheme)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.Scheme
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Scheme) (*model.Scheme, error)); ok {
		return rf(scheme)
	}
	if rf, ok := ret.Get(0).(func(*model.Scheme) *model.Scheme); ok {
		r0 = rf(scheme)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Scheme)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Scheme) error); ok {
		r1 = rf(scheme)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSchemeStore creates a new instance of SchemeStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSchemeStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *SchemeStore {
	mock := &SchemeStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
