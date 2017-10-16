// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import raft "github.com/hashicorp/raft"

// RaftStore is an autogenerated mock type for the RaftStore type
type RaftStore struct {
	mock.Mock
}

// Join provides a mock function with given fields: _a0, _a1
func (_m *RaftStore) Join(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Open provides a mock function with given fields: _a0, _a1
func (_m *RaftStore) Open(_a0 bool, _a1 raft.ServerID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, raft.ServerID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
