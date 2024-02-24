package api

import (
	"errors"
	"time"

	"github.com/enbility/spine-go/model"
)

type EVChargeStateType string

const (
	EVChargeStateTypeUnknown   EVChargeStateType = "Unknown"
	EVChargeStateTypeUnplugged EVChargeStateType = "unplugged"
	EVChargeStateTypeError     EVChargeStateType = "error"
	EVChargeStateTypePaused    EVChargeStateType = "paused"
	EVChargeStateTypeActive    EVChargeStateType = "active"
	EVChargeStateTypeFinished  EVChargeStateType = "finished"
)

// Defines a phase specific limit
type LoadLimitsPhase struct {
	Phase    model.ElectricalConnectionPhaseNameType
	IsActive bool
	Value    float64
}

// identification
type IdentificationItem struct {
	// the identification value
	Value string

	// the type of the identification value, e.g.
	ValueType model.IdentificationTypeType
}

type EVChargeStrategyType string

const (
	EVChargeStrategyTypeUnknown        EVChargeStrategyType = "unknown"
	EVChargeStrategyTypeNoDemand       EVChargeStrategyType = "nodemand"
	EVChargeStrategyTypeDirectCharging EVChargeStrategyType = "directcharging"
	EVChargeStrategyTypeMinSoC         EVChargeStrategyType = "minsoc"
	EVChargeStrategyTypeTimedCharging  EVChargeStrategyType = "timedcharging"
)

// Contains details about the actual demands from the EV
//
// General:
//   - If duration and energy is 0, charge mode is EVChargeStrategyTypeNoDemand
//   - If duration is 0, charge mode is EVChargeStrategyTypeDirectCharging and the slots should cover at least 48h
//   - If both are != 0, charge mode is EVChargeStrategyTypeTimedCharging and the slots should cover at least the duration, but at max 168h (7d)
type Demand struct {
	MinDemand          float64 // minimum demand in Wh to reach the minSoC setting, 0 if not set
	OptDemand          float64 // demand in Wh to reach the timer SoC setting
	MaxDemand          float64 // the maximum possible demand until the battery is full
	DurationUntilStart float64 // the duration in s from now until charging will start, this could be in the future but usualy is now
	DurationUntilEnd   float64 // the duration in s from now until minDemand or optDemand has to be reached, 0 if direct charge strategy is active
}

// Contains details about an EV generated charging plan
type ChargePlan struct {
	Slots []ChargePlanSlotValue // Individual charging slot details
}

// Contains details about a charging plan slot
type ChargePlanSlotValue struct {
	Start    time.Time // The start time of the slot
	End      time.Time // The duration of the slot
	Value    float64   // planned power value
	MinValue float64   // minimum power value
	MaxValue float64   // maximum power value
}

// Details about the time slot constraints
type TimeSlotConstraints struct {
	MinSlots             uint          // the minimum number of slots, no minimum if 0
	MaxSlots             uint          // the maximum number of slots, unlimited if 0
	MinSlotDuration      time.Duration // the minimum duration of a slot, no minimum if 0
	MaxSlotDuration      time.Duration // the maximum duration of a slot, unlimited if 0
	SlotDurationStepSize time.Duration // the duration has to be a multiple of this value if != 0
}

// Details about the incentive slot constraints
type IncentiveSlotConstraints struct {
	MinSlots uint // the minimum number of slots, no minimum if 0
	MaxSlots uint // the maximum number of slots, unlimited if 0
}

// details about the boundary
type TierBoundaryDescription struct {
	// the id of the boundary
	Id uint

	// the type of the boundary
	Type model.TierBoundaryTypeType

	// the unit of the boundary
	Unit model.UnitOfMeasurementType
}

// details about incentive
type IncentiveDescription struct {
	// the id of the incentive
	Id uint

	// the type of the incentive
	Type model.IncentiveTypeType

	// the currency of the incentive, if it is price based
	Currency model.CurrencyType
}

// Contains about one tier in a tariff
type IncentiveTableDescriptionTier struct {
	// the id of the tier
	Id uint

	// the tiers type
	Type model.TierTypeType

	// each tear has 1 to 3 boundaries
	// used for different power limits, e.g. 0-1kW x€, 1-3kW y€, ...
	Boundaries []TierBoundaryDescription

	// each tier has 1 to 3 incentives
	//   - price/costs (absolute or relative)
	//   - renewable energy percentage
	//   - CO2 emissions
	Incentives []IncentiveDescription
}

// Contains details about a tariff
type IncentiveTariffDescription struct {
	// each tariff can have 1 to 3 tiers
	Tiers []IncentiveTableDescriptionTier
}

// Contains details about power limits or incentives for a defined timeframe
type DurationSlotValue struct {
	Duration time.Duration // Duration of this slot
	Value    float64       // Energy Cost or Power Limit
}

// value if the UCEVCC communication standard is unknown
const (
	UCEVCCCommunicationStandardUnknown string = "unknown"
)

// type for usecase specfic event names
type UseCaseEventType string

const (
	// UCCEVC

	// EV provided an energy demand
	UCCEVCEnergyDemandProvided UseCaseEventType = "ucCEVCEnergyDemandProvided"

	// EV provided a charge plan
	UCCEVCChargePlanProvided UseCaseEventType = "ucCEVCChargePlanProvided"

	// EV provided a charge plan constraints
	UCCEVCChargePlanConstraintsProvided UseCaseEventType = "ucCEVCChargePlanConstraintsProvided"

	UCCEVCIncentiveDescriptionsRequired UseCaseEventType = "ucCEVCIncentiveDescriptionsRequired"

	// EV incentive table data updated
	UCCEVCIncentiveTableDataUpdate UseCaseEventType = "ucCEVCIncentiveTableDataUpdate"

	// EV requested power limits
	UCCEVPowerLimitsRequested UseCaseEventType = "ucCEVPowerLimitsRequested"

	// EV requested incentives
	UCCEVCIncentivesRequested UseCaseEventType = "ucCEVCIncentivesRequested"

	// UCEVCC

	// An EV was connected
	//
	// Use Case EVCC, Scenario 1
	UCEVCCEventConnected UseCaseEventType = "ucEVCCEventConnected"

	// An EV was disconnected
	//
	// Use Case EVCC, Scenario 8
	UCEVCCEventDisconnected UseCaseEventType = "ucEVCCEventDisconnected"

	// EV communication standard data was updated
	//
	// Use Case EVCC, Scenario 2
	// Note: the referred data may be updated together with all other configuration items of this use case
	UCEVCCCommunicationStandardConfigurationDataUpdate UseCaseEventType = "ucEVCCCommunicationStandardConfigurationDataUpdate"

	// EV asymmetric charging data was updated
	//
	// Use Case EVCC, Scenario 3
	//
	// Note: the referred data may be updated together with all other configuration items of this use case
	UCEVCCAsymmetricChargingConfigurationDataUpdate UseCaseEventType = "ucEVCCAsymmetricChargingConfigurationDataUpdate"

	// EV identificationdata was updated
	//
	// Use Case EVCC, Scenario 4
	UCEVCCIdentificationDataUpdate UseCaseEventType = "ucEVCCIdentificationDataUpdate"

	// EV manufacturer data was updated
	//
	// Use Case EVCC, Scenario 5
	UCEVCCManufacturerDataUpdate UseCaseEventType = "ucEVCCManufacturerDataUpdate"

	// EV charging power limits
	//
	// Use Case EVCC, Scenario 6
	UCEVCCChargingPowerLimitsDataUpdate UseCaseEventType = "ucEVCCChargingPowerLimitsDataUpdate"

	// EV permitted power limits updated
	//
	// Use Case EVCC, Scenario 7
	UCEVCCSleepModeDataUpdate UseCaseEventType = "ucEVCCSleepModeDataUpdate"

	// UCEVCEM

	// EV number of connected phases data updated
	//
	// Use Case EVCEM, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVCEMNumberOfConnectedPhasesDataUpdate UseCaseEventType = "ucEVCEMNumberOfConnectedPhasesDataUpdate"

	// EV current measurement data updated
	//
	// Use Case EVCEM, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVCEMCurrentMeasurementDataUpdate UseCaseEventType = "ucEVCEMCurrentMeasurementDataUpdate"

	// EV power measurement data updated
	//
	// Use Case EVCEM, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVCEMPowerMeasurementDataUpdate UseCaseEventType = "ucEVCEMCurrentMeasurementDataUpdate"

	// EV charging energy measurement data updated
	//
	// Use Case EVCEM, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVCEMChargingEnergyMeasurementDataUpdate UseCaseEventType = "UCEVCEMChargingEnergyMeasurementDataUpdate"

	// UCEVSECC

	// An EVSE was connected
	UCEVSECCEventConnected UseCaseEventType = "ucEVSEConnected"

	// An EVSE was disconnected
	UCEVSECCEventDisconnected UseCaseEventType = "ucEVSEDisonnected"

	// EVSE manufacturer data was updated
	//
	// Use Case EVSECC, Scenario 1
	UCEVSECCEventManufacturerUpdate UseCaseEventType = "ucEVSEManufacturerUpdate"

	// EVSE operation state was updated
	//
	// Use Case EVSECC, Scenario 2
	UCEVSECCEventOperationStateUpdate UseCaseEventType = "ucEVSEOperationStateUpdate"

	// UCEVSOC

	// EV state of charge data was updated
	//
	// Use Case EVSOC, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVSOCStateOfChargeMeasurementDataUpdate UseCaseEventType = "ucEVSOCStateOfChargeMeasurementDataUpdate"

	// EV nominal capacity data was updated
	//
	// Use Case EVSOC, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVSOCNominalCapacityMeasurementDataUpdate UseCaseEventType = "ucEVSOCNominalCapacityMeasurementDataUpdate"

	// EV state of health data was updated
	//
	// Use Case EVSOC, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	EVSOCStateOfHealthMeasurementDataUpdate UseCaseEventType = "ucEVSOCStateOfHealthMeasurementDataUpdate"

	// EV actual range data was updated
	//
	// Use Case EVSOC, Scenario 4
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCEVSOCActualRangeMeasurementDataUpdate UseCaseEventType = "ucEVSOCActualRangeMeasurementDataUpdate"

	// MGCP

	// Grid momentary power consumption/production data updated
	//
	// Use Case MGCP, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMGCPPowerTotalMeasurementDataUpdate UseCaseEventType = "ucMGCPPowerTotalMeasurementDataUpdate"

	// MTotal grid feed in energy data updated
	//
	// Use Case MGCP, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMGCPGridFeedInMeasurementDataUpdate UseCaseEventType = "ucMGCPGridFeedInMeasurementDataUpdate"

	// Total grid consumed energy data updated
	//
	// Use Case MGCP, Scenario 4
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMGCPGridConsumptionMeasurementDataUpdate UseCaseEventType = "ucMGCPGridConsumptionMeasurementDataUpdate"

	// Phase specific momentary current consumption/production phase detail data updated
	//
	// Use Case MGCP, Scenario 5
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMGCPCurrentsMeasurementDataUpdate UseCaseEventType = "ucMGCPCurrentsMeasurementDataUpdate"

	// Phase specific voltage at the grid connection point
	//
	// Use Case MGCP, Scenario 6
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMGCPVoltagesMeasurementDataUpdate UseCaseEventType = "ucMGCPVoltagesMeasurementDataUpdate"

	// Grid frequency data updated
	//
	// Use Case MGCP, Scenario 7
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMGCPFrequencyMeasurementDataUpdate UseCaseEventType = "ucMGCPFrequencyMeasurementDataUpdate"

	// UCMPC

	// Total momentary active power consumption or production
	//
	// Use Case MCP, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCPowerTotalMeasurementDataUpdate UseCaseEventType = "ucMPCPowerTotalMeasurementDataUpdate"

	// Phase specific momentary active power consumption or production
	//
	// Use Case MCP, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCPowerPerPhaseMeasurementDataUpdate UseCaseEventType = "ucMPCPowerPerPhaseMeasurementDataUpdate"

	// Total energy consumed
	//
	// Use Case MCP, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCEnergyConsumedMeasurementDataUpdate UseCaseEventType = "ucMPCEnergyConsumedMeasurementDataUpdate"

	// Total energy produced
	//
	// Use Case MCP, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCEnergyProcudedMeasurementDataUpdate UseCaseEventType = "ucMPCEnergyProcudedMeasurementDataUpdate"

	// Phase specific momentary current consumption or production
	//
	// Use Case MCP, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCCurrentsMeasurementDataUpdate UseCaseEventType = "ucMPCCurrentsMeasurementDataUpdate"

	// Phase specific voltage
	//
	// Use Case MCP, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCVoltagesMeasurementDataUpdate UseCaseEventType = "ucMPCVoltagesMeasurementDataUpdate"

	// Power network frequency data updated
	//
	// Use Case MCP, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCMPCFrequencyMeasurementDataUpdate UseCaseEventType = "ucMPCFrequencyMeasurementDataUpdate"

	// UCOPEV

	// EV load control obligation limit data updated
	UCOPEVLoadControlLimitDataUpdate UseCaseEventType = "ucOPEVLoadControlLimitDataUpdate"

	// UCOSCEV

	// EV load control recommendation limit data updated
	//
	// Use Case OSCEV, Scenario 1
	UCOSCEVLoadControlLimitDataUpdate UseCaseEventType = "ucOSCEVLoadControlLimitDataUpdate"

	// UCVABD

	// Battery System (dis)charge power data updated
	//
	// Use Case VABD, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCVABDPowerTotalMeasurementDataUpdate UseCaseEventType = "ucVABDPowerTotalMeasurementDataUpdate"

	// Battery System cumulated charge energy data updated
	//
	// Use Case VABD, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCVABDChargeMeasurementDataUpdate UseCaseEventType = "ucVABDChargeMeasurementDataUpdate"

	// Battery System cumulated discharge energy data updated
	//
	// Use Case VABD, Scenario 2
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCVABDDischargeMeasurementDataUpdate UseCaseEventType = "ucVABDDischargeMeasurementDataUpdate"

	// Battery System state of charge data updated
	//
	// Use Case VABD, Scenario 4
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCVABDStateOfChargeMeasurementDataUpdate UseCaseEventType = "ucVABDStateOfChargeMeasurementDataUpdate"

	// UCVAPD

	// PV System total power data updated
	//
	// Use Case VAPD, Scenario 1
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCVAPDPowerTotalMeasurementDataUpdate UseCaseEventType = "ucVAPDPowerTotalMeasurementDataUpdate"

	// PV System nominal peak power data updated
	//
	// Use Case VAPD, Scenario 2
	UCVAPDPeakPowerDataUpdate UseCaseEventType = "ucVAPDPeakPowerDataUpdate"

	// PV System total yield data updated
	//
	// Use Case VAPD, Scenario 3
	//
	// Note: the referred data may be updated together with all other measurement items of this use case
	UCVAPDYieldTotalMeasurementDataUpdate UseCaseEventType = "ucVAPDYieldTotalMeasurementDataUpdate"
)

var ErrNoCompatibleEntity = errors.New("entity is not an compatible entity")