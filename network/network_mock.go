package network

import (
	"testing"
)

type NetworkMockIf interface {
	InitMock(t *testing.T)
	ExpectGetRequestToCall(string, []byte)
	ExpectedCallsCalled()
}

type NetworkMock struct {
	t             *testing.T
	expectedCalls []getRequestCall
}

type getRequestCall struct {
	called  bool
	arg     string
	returns []byte
}

func (n *NetworkMock) GetRequest(url string) []byte {
	for i, elem := range n.expectedCalls {
		if !elem.called && elem.arg == url {
			elem.called = true
			n.expectedCalls[i] = elem
			return elem.returns
		}
	}
	n.t.Fatalf("Unexpected GetRequest was called with: [%s]", url)
	return []byte{}
}

func (n *NetworkMock) InitMock(t *testing.T) {
	n.t = t
}

func (n *NetworkMock) ExpectGetRequestToCall(arg string, ret []byte) {
	n.expectedCalls = append(n.expectedCalls, getRequestCall{false, arg, ret})
}

func (n *NetworkMock) ExpectedCallsCalled() {
	for _, elem := range n.expectedCalls {
		if !elem.called {
			n.t.Fatalf("Expected GetRequest was not called")
		}
	}
}
