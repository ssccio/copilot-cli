// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/cli/deploy/workload.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	addon "github.com/aws/copilot-cli/internal/pkg/addon"
	cloudformation "github.com/aws/copilot-cli/internal/pkg/aws/cloudformation"
	cloudformation0 "github.com/aws/copilot-cli/internal/pkg/deploy/cloudformation"
	dockerengine "github.com/aws/copilot-cli/internal/pkg/docker/dockerengine"
	gomock "github.com/golang/mock/gomock"
)

// MockActionRecommender is a mock of ActionRecommender interface.
type MockActionRecommender struct {
	ctrl     *gomock.Controller
	recorder *MockActionRecommenderMockRecorder
}

// MockActionRecommenderMockRecorder is the mock recorder for MockActionRecommender.
type MockActionRecommenderMockRecorder struct {
	mock *MockActionRecommender
}

// NewMockActionRecommender creates a new mock instance.
func NewMockActionRecommender(ctrl *gomock.Controller) *MockActionRecommender {
	mock := &MockActionRecommender{ctrl: ctrl}
	mock.recorder = &MockActionRecommenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionRecommender) EXPECT() *MockActionRecommenderMockRecorder {
	return m.recorder
}

// RecommendedActions mocks base method.
func (m *MockActionRecommender) RecommendedActions() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecommendedActions")
	ret0, _ := ret[0].([]string)
	return ret0
}

// RecommendedActions indicates an expected call of RecommendedActions.
func (mr *MockActionRecommenderMockRecorder) RecommendedActions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecommendedActions", reflect.TypeOf((*MockActionRecommender)(nil).RecommendedActions))
}

// MockrepositoryService is a mock of repositoryService interface.
type MockrepositoryService struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryServiceMockRecorder
}

// MockrepositoryServiceMockRecorder is the mock recorder for MockrepositoryService.
type MockrepositoryServiceMockRecorder struct {
	mock *MockrepositoryService
}

// NewMockrepositoryService creates a new mock instance.
func NewMockrepositoryService(ctrl *gomock.Controller) *MockrepositoryService {
	mock := &MockrepositoryService{ctrl: ctrl}
	mock.recorder = &MockrepositoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockrepositoryService) EXPECT() *MockrepositoryServiceMockRecorder {
	return m.recorder
}

// BuildAndPush mocks base method.
func (m *MockrepositoryService) BuildAndPush(ctx context.Context, args *dockerengine.BuildArguments, w io.Writer) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAndPush", ctx, args, w)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildAndPush indicates an expected call of BuildAndPush.
func (mr *MockrepositoryServiceMockRecorder) BuildAndPush(ctx, args, w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndPush", reflect.TypeOf((*MockrepositoryService)(nil).BuildAndPush), ctx, args, w)
}

// Login mocks base method.
func (m *MockrepositoryService) Login() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockrepositoryServiceMockRecorder) Login() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockrepositoryService)(nil).Login))
}

// Mocktemplater is a mock of templater interface.
type Mocktemplater struct {
	ctrl     *gomock.Controller
	recorder *MocktemplaterMockRecorder
}

// MocktemplaterMockRecorder is the mock recorder for Mocktemplater.
type MocktemplaterMockRecorder struct {
	mock *Mocktemplater
}

// NewMocktemplater creates a new mock instance.
func NewMocktemplater(ctrl *gomock.Controller) *Mocktemplater {
	mock := &Mocktemplater{ctrl: ctrl}
	mock.recorder = &MocktemplaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktemplater) EXPECT() *MocktemplaterMockRecorder {
	return m.recorder
}

// Template mocks base method.
func (m *Mocktemplater) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MocktemplaterMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*Mocktemplater)(nil).Template))
}

// MockstackBuilder is a mock of stackBuilder interface.
type MockstackBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockstackBuilderMockRecorder
}

// MockstackBuilderMockRecorder is the mock recorder for MockstackBuilder.
type MockstackBuilderMockRecorder struct {
	mock *MockstackBuilder
}

// NewMockstackBuilder creates a new mock instance.
func NewMockstackBuilder(ctrl *gomock.Controller) *MockstackBuilder {
	mock := &MockstackBuilder{ctrl: ctrl}
	mock.recorder = &MockstackBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstackBuilder) EXPECT() *MockstackBuilderMockRecorder {
	return m.recorder
}

// Package mocks base method.
func (m *MockstackBuilder) Package(arg0 addon.PackageConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Package", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Package indicates an expected call of Package.
func (mr *MockstackBuilderMockRecorder) Package(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Package", reflect.TypeOf((*MockstackBuilder)(nil).Package), arg0)
}

// Parameters mocks base method.
func (m *MockstackBuilder) Parameters() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parameters indicates an expected call of Parameters.
func (mr *MockstackBuilderMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MockstackBuilder)(nil).Parameters))
}

// Template mocks base method.
func (m *MockstackBuilder) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockstackBuilderMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockstackBuilder)(nil).Template))
}

// MockstackSerializer is a mock of stackSerializer interface.
type MockstackSerializer struct {
	ctrl     *gomock.Controller
	recorder *MockstackSerializerMockRecorder
}

// MockstackSerializerMockRecorder is the mock recorder for MockstackSerializer.
type MockstackSerializerMockRecorder struct {
	mock *MockstackSerializer
}

// NewMockstackSerializer creates a new mock instance.
func NewMockstackSerializer(ctrl *gomock.Controller) *MockstackSerializer {
	mock := &MockstackSerializer{ctrl: ctrl}
	mock.recorder = &MockstackSerializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstackSerializer) EXPECT() *MockstackSerializerMockRecorder {
	return m.recorder
}

// SerializedParameters mocks base method.
func (m *MockstackSerializer) SerializedParameters() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SerializedParameters")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SerializedParameters indicates an expected call of SerializedParameters.
func (mr *MockstackSerializerMockRecorder) SerializedParameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SerializedParameters", reflect.TypeOf((*MockstackSerializer)(nil).SerializedParameters))
}

// Template mocks base method.
func (m *MockstackSerializer) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockstackSerializerMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockstackSerializer)(nil).Template))
}

// MockendpointGetter is a mock of endpointGetter interface.
type MockendpointGetter struct {
	ctrl     *gomock.Controller
	recorder *MockendpointGetterMockRecorder
}

// MockendpointGetterMockRecorder is the mock recorder for MockendpointGetter.
type MockendpointGetterMockRecorder struct {
	mock *MockendpointGetter
}

// NewMockendpointGetter creates a new mock instance.
func NewMockendpointGetter(ctrl *gomock.Controller) *MockendpointGetter {
	mock := &MockendpointGetter{ctrl: ctrl}
	mock.recorder = &MockendpointGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockendpointGetter) EXPECT() *MockendpointGetterMockRecorder {
	return m.recorder
}

// ServiceDiscoveryEndpoint mocks base method.
func (m *MockendpointGetter) ServiceDiscoveryEndpoint() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceDiscoveryEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceDiscoveryEndpoint indicates an expected call of ServiceDiscoveryEndpoint.
func (mr *MockendpointGetterMockRecorder) ServiceDiscoveryEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceDiscoveryEndpoint", reflect.TypeOf((*MockendpointGetter)(nil).ServiceDiscoveryEndpoint))
}

// MockserviceDeployer is a mock of serviceDeployer interface.
type MockserviceDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockserviceDeployerMockRecorder
}

// MockserviceDeployerMockRecorder is the mock recorder for MockserviceDeployer.
type MockserviceDeployerMockRecorder struct {
	mock *MockserviceDeployer
}

// NewMockserviceDeployer creates a new mock instance.
func NewMockserviceDeployer(ctrl *gomock.Controller) *MockserviceDeployer {
	mock := &MockserviceDeployer{ctrl: ctrl}
	mock.recorder = &MockserviceDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserviceDeployer) EXPECT() *MockserviceDeployerMockRecorder {
	return m.recorder
}

// DeployService mocks base method.
func (m *MockserviceDeployer) DeployService(conf cloudformation0.StackConfiguration, bucketName string, opts ...cloudformation.StackOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{conf, bucketName}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeployService", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployService indicates an expected call of DeployService.
func (mr *MockserviceDeployerMockRecorder) DeployService(conf, bucketName interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{conf, bucketName}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployService", reflect.TypeOf((*MockserviceDeployer)(nil).DeployService), varargs...)
}

// MockdeployedTemplateGetter is a mock of deployedTemplateGetter interface.
type MockdeployedTemplateGetter struct {
	ctrl     *gomock.Controller
	recorder *MockdeployedTemplateGetterMockRecorder
}

// MockdeployedTemplateGetterMockRecorder is the mock recorder for MockdeployedTemplateGetter.
type MockdeployedTemplateGetterMockRecorder struct {
	mock *MockdeployedTemplateGetter
}

// NewMockdeployedTemplateGetter creates a new mock instance.
func NewMockdeployedTemplateGetter(ctrl *gomock.Controller) *MockdeployedTemplateGetter {
	mock := &MockdeployedTemplateGetter{ctrl: ctrl}
	mock.recorder = &MockdeployedTemplateGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdeployedTemplateGetter) EXPECT() *MockdeployedTemplateGetterMockRecorder {
	return m.recorder
}

// Template mocks base method.
func (m *MockdeployedTemplateGetter) Template(stackName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template", stackName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockdeployedTemplateGetterMockRecorder) Template(stackName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockdeployedTemplateGetter)(nil).Template), stackName)
}

// Mockspinner is a mock of spinner interface.
type Mockspinner struct {
	ctrl     *gomock.Controller
	recorder *MockspinnerMockRecorder
}

// MockspinnerMockRecorder is the mock recorder for Mockspinner.
type MockspinnerMockRecorder struct {
	mock *Mockspinner
}

// NewMockspinner creates a new mock instance.
func NewMockspinner(ctrl *gomock.Controller) *Mockspinner {
	mock := &Mockspinner{ctrl: ctrl}
	mock.recorder = &MockspinnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockspinner) EXPECT() *MockspinnerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *Mockspinner) Start(label string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", label)
}

// Start indicates an expected call of Start.
func (mr *MockspinnerMockRecorder) Start(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockspinner)(nil).Start), label)
}

// Stop mocks base method.
func (m *Mockspinner) Stop(label string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", label)
}

// Stop indicates an expected call of Stop.
func (mr *MockspinnerMockRecorder) Stop(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Mockspinner)(nil).Stop), label)
}

// MocklabeledTermPrinter is a mock of labeledTermPrinter interface.
type MocklabeledTermPrinter struct {
	ctrl     *gomock.Controller
	recorder *MocklabeledTermPrinterMockRecorder
}

// MocklabeledTermPrinterMockRecorder is the mock recorder for MocklabeledTermPrinter.
type MocklabeledTermPrinterMockRecorder struct {
	mock *MocklabeledTermPrinter
}

// NewMocklabeledTermPrinter creates a new mock instance.
func NewMocklabeledTermPrinter(ctrl *gomock.Controller) *MocklabeledTermPrinter {
	mock := &MocklabeledTermPrinter{ctrl: ctrl}
	mock.recorder = &MocklabeledTermPrinterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocklabeledTermPrinter) EXPECT() *MocklabeledTermPrinterMockRecorder {
	return m.recorder
}

// IsDone mocks base method.
func (m *MocklabeledTermPrinter) IsDone() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDone")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDone indicates an expected call of IsDone.
func (mr *MocklabeledTermPrinterMockRecorder) IsDone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDone", reflect.TypeOf((*MocklabeledTermPrinter)(nil).IsDone))
}

// Print mocks base method.
func (m *MocklabeledTermPrinter) Print() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Print")
}

// Print indicates an expected call of Print.
func (mr *MocklabeledTermPrinterMockRecorder) Print() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MocklabeledTermPrinter)(nil).Print))
}

// MockdockerEngineRunChecker is a mock of dockerEngineRunChecker interface.
type MockdockerEngineRunChecker struct {
	ctrl     *gomock.Controller
	recorder *MockdockerEngineRunCheckerMockRecorder
}

// MockdockerEngineRunCheckerMockRecorder is the mock recorder for MockdockerEngineRunChecker.
type MockdockerEngineRunCheckerMockRecorder struct {
	mock *MockdockerEngineRunChecker
}

// NewMockdockerEngineRunChecker creates a new mock instance.
func NewMockdockerEngineRunChecker(ctrl *gomock.Controller) *MockdockerEngineRunChecker {
	mock := &MockdockerEngineRunChecker{ctrl: ctrl}
	mock.recorder = &MockdockerEngineRunCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdockerEngineRunChecker) EXPECT() *MockdockerEngineRunCheckerMockRecorder {
	return m.recorder
}

// CheckDockerEngineRunning mocks base method.
func (m *MockdockerEngineRunChecker) CheckDockerEngineRunning() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDockerEngineRunning")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDockerEngineRunning indicates an expected call of CheckDockerEngineRunning.
func (mr *MockdockerEngineRunCheckerMockRecorder) CheckDockerEngineRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDockerEngineRunning", reflect.TypeOf((*MockdockerEngineRunChecker)(nil).CheckDockerEngineRunning))
}

// MocktimeoutError is a mock of timeoutError interface.
type MocktimeoutError struct {
	ctrl     *gomock.Controller
	recorder *MocktimeoutErrorMockRecorder
}

// MocktimeoutErrorMockRecorder is the mock recorder for MocktimeoutError.
type MocktimeoutErrorMockRecorder struct {
	mock *MocktimeoutError
}

// NewMocktimeoutError creates a new mock instance.
func NewMocktimeoutError(ctrl *gomock.Controller) *MocktimeoutError {
	mock := &MocktimeoutError{ctrl: ctrl}
	mock.recorder = &MocktimeoutErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktimeoutError) EXPECT() *MocktimeoutErrorMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MocktimeoutError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MocktimeoutErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MocktimeoutError)(nil).Error))
}

// Timeout mocks base method.
func (m *MocktimeoutError) Timeout() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MocktimeoutErrorMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MocktimeoutError)(nil).Timeout))
}
