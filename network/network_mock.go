package network

import (
	"mock"
	"testing"
)

type NetworkMock struct {
	mockHelper *mock.MockHelper
}

func (n *NetworkMock) GetRequest(url string) []byte {
	n.mockHelper.AddCall("GetRequest", url)
	return []byte{}
}

func (n *NetworkMock) InitMock() {
	n.mockHelper = &mock.MockHelper{}
	n.mockHelper.InitMock()
}

func (n *NetworkMock) ExpectCall(funcName string, args ...interface{}) {
	n.mockHelper.ExpectCall(funcName, args)
}

func (n *NetworkMock) ExpectedCallsCalled(t *testing.T) {
	n.mockHelper.ExpectedCallsCalled(t)
}
