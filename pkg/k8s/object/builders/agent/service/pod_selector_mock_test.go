// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/k8s/object/transformations/pod_selector.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPodSelectorLabelGenerator is a mock of PodSelectorLabelGenerator interface.
type MockPodSelectorLabelGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockPodSelectorLabelGeneratorMockRecorder
}

// MockPodSelectorLabelGeneratorMockRecorder is the mock recorder for MockPodSelectorLabelGenerator.
type MockPodSelectorLabelGeneratorMockRecorder struct {
	mock *MockPodSelectorLabelGenerator
}

// NewMockPodSelectorLabelGenerator creates a new mock instance.
func NewMockPodSelectorLabelGenerator(ctrl *gomock.Controller) *MockPodSelectorLabelGenerator {
	mock := &MockPodSelectorLabelGenerator{ctrl: ctrl}
	mock.recorder = &MockPodSelectorLabelGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodSelectorLabelGenerator) EXPECT() *MockPodSelectorLabelGeneratorMockRecorder {
	return m.recorder
}

// GetPodLabels mocks base method.
func (m *MockPodSelectorLabelGenerator) GetPodLabels(userLabels map[string]string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodLabels", userLabels)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetPodLabels indicates an expected call of GetPodLabels.
func (mr *MockPodSelectorLabelGeneratorMockRecorder) GetPodLabels(userLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodLabels", reflect.TypeOf((*MockPodSelectorLabelGenerator)(nil).GetPodLabels), userLabels)
}

// GetPodSelectorLabels mocks base method.
func (m *MockPodSelectorLabelGenerator) GetPodSelectorLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodSelectorLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetPodSelectorLabels indicates an expected call of GetPodSelectorLabels.
func (mr *MockPodSelectorLabelGeneratorMockRecorder) GetPodSelectorLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodSelectorLabels", reflect.TypeOf((*MockPodSelectorLabelGenerator)(nil).GetPodSelectorLabels))
}
