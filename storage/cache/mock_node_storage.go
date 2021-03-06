// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian/storage/cache (interfaces: NodeStorage)

package cache

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/google/trillian/storage"
	storagepb "github.com/google/trillian/storage/storagepb"
)

// MockNodeStorage is a mock of NodeStorage interface
type MockNodeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockNodeStorageMockRecorder
}

// MockNodeStorageMockRecorder is the mock recorder for MockNodeStorage
type MockNodeStorageMockRecorder struct {
	mock *MockNodeStorage
}

// NewMockNodeStorage creates a new mock instance
func NewMockNodeStorage(ctrl *gomock.Controller) *MockNodeStorage {
	mock := &MockNodeStorage{ctrl: ctrl}
	mock.recorder = &MockNodeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNodeStorage) EXPECT() *MockNodeStorageMockRecorder {
	return _m.recorder
}

// GetSubtree mocks base method
func (_m *MockNodeStorage) GetSubtree(_param0 storage.NodeID) (*storagepb.SubtreeProto, error) {
	ret := _m.ctrl.Call(_m, "GetSubtree", _param0)
	ret0, _ := ret[0].(*storagepb.SubtreeProto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubtree indicates an expected call of GetSubtree
func (_mr *MockNodeStorageMockRecorder) GetSubtree(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubtree", arg0)
}

// SetSubtrees mocks base method
func (_m *MockNodeStorage) SetSubtrees(_param0 []*storagepb.SubtreeProto) error {
	ret := _m.ctrl.Call(_m, "SetSubtrees", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSubtrees indicates an expected call of SetSubtrees
func (_mr *MockNodeStorageMockRecorder) SetSubtrees(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSubtrees", arg0)
}
