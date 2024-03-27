package ucevcc

import (
	"github.com/enbility/cemd/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

//go:generate mockery

// interface for the EV Commissioning and Configuration UseCase
type UCEVCCInterface interface {
	api.UseCaseInterface

	// return the current charge state of the EV
	//
	// parameters:
	//   - entity: the entity of the EV
	ChargeState(entity spineapi.EntityRemoteInterface) (api.EVChargeStateType, error)

	// Scenario 1 & 8

	// return if the EV is connected
	//
	// parameters:
	//   - entity: the entity of the EV
	EVConnected(entity spineapi.EntityRemoteInterface) bool

	// Scenario 2

	// return the current communication standard type used to communicate between EVSE and EV
	//
	// parameters:
	//   - entity: the entity of the EV
	CommunicationStandard(entity spineapi.EntityRemoteInterface) (model.DeviceConfigurationKeyValueStringType, error)

	// Scenario 3

	// return if the EV supports asymmetric charging
	//
	// parameters:
	//   - entity: the entity of the EV
	AsymmetricChargingSupport(entity spineapi.EntityRemoteInterface) (bool, error)

	// Scenario 4

	// return the identifications of the currently connected EV or nil if not available
	// these can be multiple, e.g. PCID, Mac Address, RFID
	//
	// parameters:
	//   - entity: the entity of the EV
	Identifications(entity spineapi.EntityRemoteInterface) ([]api.IdentificationItem, error)

	// Scenario 5

	// the manufacturer data of an EVSE
	// returns deviceName, serialNumber, error
	//
	// parameters:
	//   - entity: the entity of the EV
	ManufacturerData(entity spineapi.EntityRemoteInterface) (api.ManufacturerData, error)

	// Scenario 6

	// return the minimum, maximum charging and, standby power of the connected EV
	//
	// parameters:
	//   - entity: the entity of the EV
	CurrentLimits(entity spineapi.EntityRemoteInterface) (float64, float64, float64, error)

	// Scenario 7

	// is the EV in sleep mode
	//
	// parameters:
	//   - entity: the entity of the EV
	IsInSleepMode(entity spineapi.EntityRemoteInterface) (bool, error)
}
