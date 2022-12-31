// Code generated by MockGen. DO NOT EDIT.
// Source: emobility.go

// Package emobility is a generated GoMock package.
package emobility

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmobilityDataProvider is a mock of EmobilityDataProvider interface.
type MockEmobilityDataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockEmobilityDataProviderMockRecorder
}

// MockEmobilityDataProviderMockRecorder is the mock recorder for MockEmobilityDataProvider.
type MockEmobilityDataProviderMockRecorder struct {
	mock *MockEmobilityDataProvider
}

// NewMockEmobilityDataProvider creates a new mock instance.
func NewMockEmobilityDataProvider(ctrl *gomock.Controller) *MockEmobilityDataProvider {
	mock := &MockEmobilityDataProvider{ctrl: ctrl}
	mock.recorder = &MockEmobilityDataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmobilityDataProvider) EXPECT() *MockEmobilityDataProviderMockRecorder {
	return m.recorder
}

// EVProvidedChargePlan mocks base method.
func (m *MockEmobilityDataProvider) EVProvidedChargePlan(data []EVDurationSlotValue) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EVProvidedChargePlan", data)
}

// EVProvidedChargePlan indicates an expected call of EVProvidedChargePlan.
func (mr *MockEmobilityDataProviderMockRecorder) EVProvidedChargePlan(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVProvidedChargePlan", reflect.TypeOf((*MockEmobilityDataProvider)(nil).EVProvidedChargePlan), data)
}

// EVProvidedChargeStrategy mocks base method.
func (m *MockEmobilityDataProvider) EVProvidedChargeStrategy(strategy EVChargeStrategyType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EVProvidedChargeStrategy", strategy)
}

// EVProvidedChargeStrategy indicates an expected call of EVProvidedChargeStrategy.
func (mr *MockEmobilityDataProviderMockRecorder) EVProvidedChargeStrategy(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVProvidedChargeStrategy", reflect.TypeOf((*MockEmobilityDataProvider)(nil).EVProvidedChargeStrategy), strategy)
}

// EVRequestIncentives mocks base method.
func (m *MockEmobilityDataProvider) EVRequestIncentives(demand EVDemand, constraints EVIncentiveSlotConstraints) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EVRequestIncentives", demand, constraints)
}

// EVRequestIncentives indicates an expected call of EVRequestIncentives.
func (mr *MockEmobilityDataProviderMockRecorder) EVRequestIncentives(demand, constraints interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVRequestIncentives", reflect.TypeOf((*MockEmobilityDataProvider)(nil).EVRequestIncentives), demand, constraints)
}

// EVRequestPowerLimits mocks base method.
func (m *MockEmobilityDataProvider) EVRequestPowerLimits(demand EVDemand, constraints EVTimeSlotConstraints) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EVRequestPowerLimits", demand, constraints)
}

// EVRequestPowerLimits indicates an expected call of EVRequestPowerLimits.
func (mr *MockEmobilityDataProviderMockRecorder) EVRequestPowerLimits(demand, constraints interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVRequestPowerLimits", reflect.TypeOf((*MockEmobilityDataProvider)(nil).EVRequestPowerLimits), demand, constraints)
}

// MockEmobilityI is a mock of EmobilityI interface.
type MockEmobilityI struct {
	ctrl     *gomock.Controller
	recorder *MockEmobilityIMockRecorder
}

// MockEmobilityIMockRecorder is the mock recorder for MockEmobilityI.
type MockEmobilityIMockRecorder struct {
	mock *MockEmobilityI
}

// NewMockEmobilityI creates a new mock instance.
func NewMockEmobilityI(ctrl *gomock.Controller) *MockEmobilityI {
	mock := &MockEmobilityI{ctrl: ctrl}
	mock.recorder = &MockEmobilityIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmobilityI) EXPECT() *MockEmobilityIMockRecorder {
	return m.recorder
}

// EVChargeStrategy mocks base method.
func (m *MockEmobilityI) EVChargeStrategy() EVChargeStrategyType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVChargeStrategy")
	ret0, _ := ret[0].(EVChargeStrategyType)
	return ret0
}

// EVChargeStrategy indicates an expected call of EVChargeStrategy.
func (mr *MockEmobilityIMockRecorder) EVChargeStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVChargeStrategy", reflect.TypeOf((*MockEmobilityI)(nil).EVChargeStrategy))
}

// EVChargedEnergy mocks base method.
func (m *MockEmobilityI) EVChargedEnergy() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVChargedEnergy")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVChargedEnergy indicates an expected call of EVChargedEnergy.
func (mr *MockEmobilityIMockRecorder) EVChargedEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVChargedEnergy", reflect.TypeOf((*MockEmobilityI)(nil).EVChargedEnergy))
}

// EVCommunicationStandard mocks base method.
func (m *MockEmobilityI) EVCommunicationStandard() (EVCommunicationStandardType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVCommunicationStandard")
	ret0, _ := ret[0].(EVCommunicationStandardType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVCommunicationStandard indicates an expected call of EVCommunicationStandard.
func (mr *MockEmobilityIMockRecorder) EVCommunicationStandard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVCommunicationStandard", reflect.TypeOf((*MockEmobilityI)(nil).EVCommunicationStandard))
}

// EVConnectedPhases mocks base method.
func (m *MockEmobilityI) EVConnectedPhases() (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVConnectedPhases")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVConnectedPhases indicates an expected call of EVConnectedPhases.
func (mr *MockEmobilityIMockRecorder) EVConnectedPhases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVConnectedPhases", reflect.TypeOf((*MockEmobilityI)(nil).EVConnectedPhases))
}

// EVCoordinatedChargingSupported mocks base method.
func (m *MockEmobilityI) EVCoordinatedChargingSupported() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVCoordinatedChargingSupported")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVCoordinatedChargingSupported indicates an expected call of EVCoordinatedChargingSupported.
func (mr *MockEmobilityIMockRecorder) EVCoordinatedChargingSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVCoordinatedChargingSupported", reflect.TypeOf((*MockEmobilityI)(nil).EVCoordinatedChargingSupported))
}

// EVCurrentChargeState mocks base method.
func (m *MockEmobilityI) EVCurrentChargeState() (EVChargeStateType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVCurrentChargeState")
	ret0, _ := ret[0].(EVChargeStateType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVCurrentChargeState indicates an expected call of EVCurrentChargeState.
func (mr *MockEmobilityIMockRecorder) EVCurrentChargeState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVCurrentChargeState", reflect.TypeOf((*MockEmobilityI)(nil).EVCurrentChargeState))
}

// EVCurrentLimits mocks base method.
func (m *MockEmobilityI) EVCurrentLimits() ([]float64, []float64, []float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVCurrentLimits")
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].([]float64)
	ret2, _ := ret[2].([]float64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EVCurrentLimits indicates an expected call of EVCurrentLimits.
func (mr *MockEmobilityIMockRecorder) EVCurrentLimits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVCurrentLimits", reflect.TypeOf((*MockEmobilityI)(nil).EVCurrentLimits))
}

// EVCurrentsPerPhase mocks base method.
func (m *MockEmobilityI) EVCurrentsPerPhase() ([]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVCurrentsPerPhase")
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVCurrentsPerPhase indicates an expected call of EVCurrentsPerPhase.
func (mr *MockEmobilityIMockRecorder) EVCurrentsPerPhase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVCurrentsPerPhase", reflect.TypeOf((*MockEmobilityI)(nil).EVCurrentsPerPhase))
}

// EVEnergyDemand mocks base method.
func (m *MockEmobilityI) EVEnergyDemand() (EVDemand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVEnergyDemand")
	ret0, _ := ret[0].(EVDemand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVEnergyDemand indicates an expected call of EVEnergyDemand.
func (mr *MockEmobilityIMockRecorder) EVEnergyDemand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVEnergyDemand", reflect.TypeOf((*MockEmobilityI)(nil).EVEnergyDemand))
}

// EVGetIncentiveConstraints mocks base method.
func (m *MockEmobilityI) EVGetIncentiveConstraints() EVIncentiveSlotConstraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVGetIncentiveConstraints")
	ret0, _ := ret[0].(EVIncentiveSlotConstraints)
	return ret0
}

// EVGetIncentiveConstraints indicates an expected call of EVGetIncentiveConstraints.
func (mr *MockEmobilityIMockRecorder) EVGetIncentiveConstraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVGetIncentiveConstraints", reflect.TypeOf((*MockEmobilityI)(nil).EVGetIncentiveConstraints))
}

// EVGetPowerConstraints mocks base method.
func (m *MockEmobilityI) EVGetPowerConstraints() EVTimeSlotConstraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVGetPowerConstraints")
	ret0, _ := ret[0].(EVTimeSlotConstraints)
	return ret0
}

// EVGetPowerConstraints indicates an expected call of EVGetPowerConstraints.
func (mr *MockEmobilityIMockRecorder) EVGetPowerConstraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVGetPowerConstraints", reflect.TypeOf((*MockEmobilityI)(nil).EVGetPowerConstraints))
}

// EVIdentification mocks base method.
func (m *MockEmobilityI) EVIdentification() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVIdentification")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVIdentification indicates an expected call of EVIdentification.
func (mr *MockEmobilityIMockRecorder) EVIdentification() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVIdentification", reflect.TypeOf((*MockEmobilityI)(nil).EVIdentification))
}

// EVOptimizationOfSelfConsumptionSupported mocks base method.
func (m *MockEmobilityI) EVOptimizationOfSelfConsumptionSupported() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVOptimizationOfSelfConsumptionSupported")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVOptimizationOfSelfConsumptionSupported indicates an expected call of EVOptimizationOfSelfConsumptionSupported.
func (mr *MockEmobilityIMockRecorder) EVOptimizationOfSelfConsumptionSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVOptimizationOfSelfConsumptionSupported", reflect.TypeOf((*MockEmobilityI)(nil).EVOptimizationOfSelfConsumptionSupported))
}

// EVPowerPerPhase mocks base method.
func (m *MockEmobilityI) EVPowerPerPhase() ([]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVPowerPerPhase")
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVPowerPerPhase indicates an expected call of EVPowerPerPhase.
func (mr *MockEmobilityIMockRecorder) EVPowerPerPhase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVPowerPerPhase", reflect.TypeOf((*MockEmobilityI)(nil).EVPowerPerPhase))
}

// EVSoC mocks base method.
func (m *MockEmobilityI) EVSoC() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVSoC")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVSoC indicates an expected call of EVSoC.
func (mr *MockEmobilityIMockRecorder) EVSoC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVSoC", reflect.TypeOf((*MockEmobilityI)(nil).EVSoC))
}

// EVSoCSupported mocks base method.
func (m *MockEmobilityI) EVSoCSupported() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVSoCSupported")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EVSoCSupported indicates an expected call of EVSoCSupported.
func (mr *MockEmobilityIMockRecorder) EVSoCSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVSoCSupported", reflect.TypeOf((*MockEmobilityI)(nil).EVSoCSupported))
}

// EVWriteIncentives mocks base method.
func (m *MockEmobilityI) EVWriteIncentives(data []EVDurationSlotValue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVWriteIncentives", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// EVWriteIncentives indicates an expected call of EVWriteIncentives.
func (mr *MockEmobilityIMockRecorder) EVWriteIncentives(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVWriteIncentives", reflect.TypeOf((*MockEmobilityI)(nil).EVWriteIncentives), data)
}

// EVWriteLoadControlLimits mocks base method.
func (m *MockEmobilityI) EVWriteLoadControlLimits(obligations, recommendations []float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVWriteLoadControlLimits", obligations, recommendations)
	ret0, _ := ret[0].(error)
	return ret0
}

// EVWriteLoadControlLimits indicates an expected call of EVWriteLoadControlLimits.
func (mr *MockEmobilityIMockRecorder) EVWriteLoadControlLimits(obligations, recommendations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVWriteLoadControlLimits", reflect.TypeOf((*MockEmobilityI)(nil).EVWriteLoadControlLimits), obligations, recommendations)
}

// EVWritePowerLimits mocks base method.
func (m *MockEmobilityI) EVWritePowerLimits(data []EVDurationSlotValue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVWritePowerLimits", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// EVWritePowerLimits indicates an expected call of EVWritePowerLimits.
func (mr *MockEmobilityIMockRecorder) EVWritePowerLimits(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVWritePowerLimits", reflect.TypeOf((*MockEmobilityI)(nil).EVWritePowerLimits), data)
}
