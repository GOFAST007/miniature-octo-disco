// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy/env.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	cloudformation0 "github.com/aws/copilot-cli/internal/pkg/aws/cloudformation"
	elbv2 "github.com/aws/copilot-cli/internal/pkg/aws/elbv2"
	config "github.com/aws/copilot-cli/internal/pkg/config"
	cloudformation1 "github.com/aws/copilot-cli/internal/pkg/deploy/cloudformation"
	stack "github.com/aws/copilot-cli/internal/pkg/deploy/cloudformation/stack"
	stack0 "github.com/aws/copilot-cli/internal/pkg/describe/stack"
	gomock "github.com/golang/mock/gomock"
)

// MockWorkspaceAddonsReaderPathGetter is a mock of WorkspaceAddonsReaderPathGetter interface.
type MockWorkspaceAddonsReaderPathGetter struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceAddonsReaderPathGetterMockRecorder
}

// MockWorkspaceAddonsReaderPathGetterMockRecorder is the mock recorder for MockWorkspaceAddonsReaderPathGetter.
type MockWorkspaceAddonsReaderPathGetterMockRecorder struct {
	mock *MockWorkspaceAddonsReaderPathGetter
}

// NewMockWorkspaceAddonsReaderPathGetter creates a new mock instance.
func NewMockWorkspaceAddonsReaderPathGetter(ctrl *gomock.Controller) *MockWorkspaceAddonsReaderPathGetter {
	mock := &MockWorkspaceAddonsReaderPathGetter{ctrl: ctrl}
	mock.recorder = &MockWorkspaceAddonsReaderPathGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceAddonsReaderPathGetter) EXPECT() *MockWorkspaceAddonsReaderPathGetterMockRecorder {
	return m.recorder
}

// EnvAddonFileAbsPath mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) EnvAddonFileAbsPath(fName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvAddonFileAbsPath", fName)
	ret0, _ := ret[0].(string)
	return ret0
}

// EnvAddonFileAbsPath indicates an expected call of EnvAddonFileAbsPath.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) EnvAddonFileAbsPath(fName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvAddonFileAbsPath", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).EnvAddonFileAbsPath), fName)
}

// EnvAddonsAbsPath mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) EnvAddonsAbsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvAddonsAbsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// EnvAddonsAbsPath indicates an expected call of EnvAddonsAbsPath.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) EnvAddonsAbsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvAddonsAbsPath", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).EnvAddonsAbsPath))
}

// ListFiles mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) ListFiles(dirPath string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", dirPath)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) ListFiles(dirPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).ListFiles), dirPath)
}

// Path mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).Path))
}

// ReadFile mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) ReadFile(fPath string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", fPath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) ReadFile(fPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).ReadFile), fPath)
}

// WorkloadAddonFileAbsPath mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) WorkloadAddonFileAbsPath(wkldName, fName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkloadAddonFileAbsPath", wkldName, fName)
	ret0, _ := ret[0].(string)
	return ret0
}

// WorkloadAddonFileAbsPath indicates an expected call of WorkloadAddonFileAbsPath.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) WorkloadAddonFileAbsPath(wkldName, fName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadAddonFileAbsPath", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).WorkloadAddonFileAbsPath), wkldName, fName)
}

// WorkloadAddonsAbsPath mocks base method.
func (m *MockWorkspaceAddonsReaderPathGetter) WorkloadAddonsAbsPath(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkloadAddonsAbsPath", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// WorkloadAddonsAbsPath indicates an expected call of WorkloadAddonsAbsPath.
func (mr *MockWorkspaceAddonsReaderPathGetterMockRecorder) WorkloadAddonsAbsPath(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadAddonsAbsPath", reflect.TypeOf((*MockWorkspaceAddonsReaderPathGetter)(nil).WorkloadAddonsAbsPath), name)
}

// MockappResourcesGetter is a mock of appResourcesGetter interface.
type MockappResourcesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockappResourcesGetterMockRecorder
}

// MockappResourcesGetterMockRecorder is the mock recorder for MockappResourcesGetter.
type MockappResourcesGetterMockRecorder struct {
	mock *MockappResourcesGetter
}

// NewMockappResourcesGetter creates a new mock instance.
func NewMockappResourcesGetter(ctrl *gomock.Controller) *MockappResourcesGetter {
	mock := &MockappResourcesGetter{ctrl: ctrl}
	mock.recorder = &MockappResourcesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockappResourcesGetter) EXPECT() *MockappResourcesGetterMockRecorder {
	return m.recorder
}

// GetAppResourcesByRegion mocks base method.
func (m *MockappResourcesGetter) GetAppResourcesByRegion(app *config.Application, region string) (*stack.AppRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppResourcesByRegion", app, region)
	ret0, _ := ret[0].(*stack.AppRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppResourcesByRegion indicates an expected call of GetAppResourcesByRegion.
func (mr *MockappResourcesGetterMockRecorder) GetAppResourcesByRegion(app, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppResourcesByRegion", reflect.TypeOf((*MockappResourcesGetter)(nil).GetAppResourcesByRegion), app, region)
}

// MockenvironmentDeployer is a mock of environmentDeployer interface.
type MockenvironmentDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockenvironmentDeployerMockRecorder
}

// MockenvironmentDeployerMockRecorder is the mock recorder for MockenvironmentDeployer.
type MockenvironmentDeployerMockRecorder struct {
	mock *MockenvironmentDeployer
}

// NewMockenvironmentDeployer creates a new mock instance.
func NewMockenvironmentDeployer(ctrl *gomock.Controller) *MockenvironmentDeployer {
	mock := &MockenvironmentDeployer{ctrl: ctrl}
	mock.recorder = &MockenvironmentDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockenvironmentDeployer) EXPECT() *MockenvironmentDeployerMockRecorder {
	return m.recorder
}

// DeployedEnvironmentParameters mocks base method.
func (m *MockenvironmentDeployer) DeployedEnvironmentParameters(app, env string) ([]*cloudformation.Parameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployedEnvironmentParameters", app, env)
	ret0, _ := ret[0].([]*cloudformation.Parameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployedEnvironmentParameters indicates an expected call of DeployedEnvironmentParameters.
func (mr *MockenvironmentDeployerMockRecorder) DeployedEnvironmentParameters(app, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployedEnvironmentParameters", reflect.TypeOf((*MockenvironmentDeployer)(nil).DeployedEnvironmentParameters), app, env)
}

// ForceUpdateOutputID mocks base method.
func (m *MockenvironmentDeployer) ForceUpdateOutputID(app, env string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUpdateOutputID", app, env)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForceUpdateOutputID indicates an expected call of ForceUpdateOutputID.
func (mr *MockenvironmentDeployerMockRecorder) ForceUpdateOutputID(app, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUpdateOutputID", reflect.TypeOf((*MockenvironmentDeployer)(nil).ForceUpdateOutputID), app, env)
}

// UpdateAndRenderEnvironment mocks base method.
func (m *MockenvironmentDeployer) UpdateAndRenderEnvironment(conf cloudformation1.StackConfiguration, bucketARN string, detach bool, opts ...cloudformation0.StackOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{conf, bucketARN, detach}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAndRenderEnvironment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAndRenderEnvironment indicates an expected call of UpdateAndRenderEnvironment.
func (mr *MockenvironmentDeployerMockRecorder) UpdateAndRenderEnvironment(conf, bucketARN, detach interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{conf, bucketARN, detach}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAndRenderEnvironment", reflect.TypeOf((*MockenvironmentDeployer)(nil).UpdateAndRenderEnvironment), varargs...)
}

// Mockpatcher is a mock of patcher interface.
type Mockpatcher struct {
	ctrl     *gomock.Controller
	recorder *MockpatcherMockRecorder
}

// MockpatcherMockRecorder is the mock recorder for Mockpatcher.
type MockpatcherMockRecorder struct {
	mock *Mockpatcher
}

// NewMockpatcher creates a new mock instance.
func NewMockpatcher(ctrl *gomock.Controller) *Mockpatcher {
	mock := &Mockpatcher{ctrl: ctrl}
	mock.recorder = &MockpatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpatcher) EXPECT() *MockpatcherMockRecorder {
	return m.recorder
}

// EnsureManagerRoleIsAllowedToUpload mocks base method.
func (m *Mockpatcher) EnsureManagerRoleIsAllowedToUpload(bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureManagerRoleIsAllowedToUpload", bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureManagerRoleIsAllowedToUpload indicates an expected call of EnsureManagerRoleIsAllowedToUpload.
func (mr *MockpatcherMockRecorder) EnsureManagerRoleIsAllowedToUpload(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureManagerRoleIsAllowedToUpload", reflect.TypeOf((*Mockpatcher)(nil).EnsureManagerRoleIsAllowedToUpload), bucketName)
}

// MockprefixListGetter is a mock of prefixListGetter interface.
type MockprefixListGetter struct {
	ctrl     *gomock.Controller
	recorder *MockprefixListGetterMockRecorder
}

// MockprefixListGetterMockRecorder is the mock recorder for MockprefixListGetter.
type MockprefixListGetterMockRecorder struct {
	mock *MockprefixListGetter
}

// NewMockprefixListGetter creates a new mock instance.
func NewMockprefixListGetter(ctrl *gomock.Controller) *MockprefixListGetter {
	mock := &MockprefixListGetter{ctrl: ctrl}
	mock.recorder = &MockprefixListGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockprefixListGetter) EXPECT() *MockprefixListGetterMockRecorder {
	return m.recorder
}

// CloudFrontManagedPrefixListID mocks base method.
func (m *MockprefixListGetter) CloudFrontManagedPrefixListID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudFrontManagedPrefixListID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudFrontManagedPrefixListID indicates an expected call of CloudFrontManagedPrefixListID.
func (mr *MockprefixListGetterMockRecorder) CloudFrontManagedPrefixListID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudFrontManagedPrefixListID", reflect.TypeOf((*MockprefixListGetter)(nil).CloudFrontManagedPrefixListID))
}

// MockenvDescriber is a mock of envDescriber interface.
type MockenvDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockenvDescriberMockRecorder
}

// MockenvDescriberMockRecorder is the mock recorder for MockenvDescriber.
type MockenvDescriberMockRecorder struct {
	mock *MockenvDescriber
}

// NewMockenvDescriber creates a new mock instance.
func NewMockenvDescriber(ctrl *gomock.Controller) *MockenvDescriber {
	mock := &MockenvDescriber{ctrl: ctrl}
	mock.recorder = &MockenvDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockenvDescriber) EXPECT() *MockenvDescriberMockRecorder {
	return m.recorder
}

// Params mocks base method.
func (m *MockenvDescriber) Params() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Params")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Params indicates an expected call of Params.
func (mr *MockenvDescriberMockRecorder) Params() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Params", reflect.TypeOf((*MockenvDescriber)(nil).Params))
}

// ValidateCFServiceDomainAliases mocks base method.
func (m *MockenvDescriber) ValidateCFServiceDomainAliases() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCFServiceDomainAliases")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateCFServiceDomainAliases indicates an expected call of ValidateCFServiceDomainAliases.
func (mr *MockenvDescriberMockRecorder) ValidateCFServiceDomainAliases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCFServiceDomainAliases", reflect.TypeOf((*MockenvDescriber)(nil).ValidateCFServiceDomainAliases))
}

// MocklbDescriber is a mock of lbDescriber interface.
type MocklbDescriber struct {
	ctrl     *gomock.Controller
	recorder *MocklbDescriberMockRecorder
}

// MocklbDescriberMockRecorder is the mock recorder for MocklbDescriber.
type MocklbDescriberMockRecorder struct {
	mock *MocklbDescriber
}

// NewMocklbDescriber creates a new mock instance.
func NewMocklbDescriber(ctrl *gomock.Controller) *MocklbDescriber {
	mock := &MocklbDescriber{ctrl: ctrl}
	mock.recorder = &MocklbDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocklbDescriber) EXPECT() *MocklbDescriberMockRecorder {
	return m.recorder
}

// DescribeRule mocks base method.
func (m *MocklbDescriber) DescribeRule(arg0 context.Context, arg1 string) (elbv2.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRule", arg0, arg1)
	ret0, _ := ret[0].(elbv2.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRule indicates an expected call of DescribeRule.
func (mr *MocklbDescriberMockRecorder) DescribeRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRule", reflect.TypeOf((*MocklbDescriber)(nil).DescribeRule), arg0, arg1)
}

// MockstackDescriber is a mock of stackDescriber interface.
type MockstackDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockstackDescriberMockRecorder
}

// MockstackDescriberMockRecorder is the mock recorder for MockstackDescriber.
type MockstackDescriberMockRecorder struct {
	mock *MockstackDescriber
}

// NewMockstackDescriber creates a new mock instance.
func NewMockstackDescriber(ctrl *gomock.Controller) *MockstackDescriber {
	mock := &MockstackDescriber{ctrl: ctrl}
	mock.recorder = &MockstackDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstackDescriber) EXPECT() *MockstackDescriberMockRecorder {
	return m.recorder
}

// Resources mocks base method.
func (m *MockstackDescriber) Resources() ([]*stack0.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resources")
	ret0, _ := ret[0].([]*stack0.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resources indicates an expected call of Resources.
func (mr *MockstackDescriberMockRecorder) Resources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resources", reflect.TypeOf((*MockstackDescriber)(nil).Resources))
}