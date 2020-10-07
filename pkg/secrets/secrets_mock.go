// Code generated by MockGen. DO NOT EDIT.
// Source: secrets.go

// Package mock_secrets is a generated GoMock package.
package secrets

import (
	secretsmanager "github.com/aws/aws-sdk-go/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSecretsManagerClient is a mock of SecretsManagerClient interface
type MockSecretsManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerClientMockRecorder
}

// MockSecretsManagerClientMockRecorder is the mock recorder for MockSecretsManagerClient
type MockSecretsManagerClientMockRecorder struct {
	mock *MockSecretsManagerClient
}

// NewMockSecretsManagerClient creates a new mock instance
func NewMockSecretsManagerClient(ctrl *gomock.Controller) *MockSecretsManagerClient {
	mock := &MockSecretsManagerClient{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretsManagerClient) EXPECT() *MockSecretsManagerClientMockRecorder {
	return m.recorder
}

// GetSecretValue mocks base method
func (m *MockSecretsManagerClient) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", input)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue
func (mr *MockSecretsManagerClientMockRecorder) GetSecretValue(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsManagerClient)(nil).GetSecretValue), input)
}
