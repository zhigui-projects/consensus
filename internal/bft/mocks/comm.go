// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import protos "github.com/SmartBFT-Go/consensus/protos"

// Comm is an autogenerated mock type for the Comm type
type Comm struct {
	mock.Mock
}

// Broadcast provides a mock function with given fields: m
func (_m *Comm) Broadcast(m *protos.Message) {
	_m.Called(m)
}