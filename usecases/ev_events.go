package usecases

import (
	"fmt"

	"github.com/DerAndereAndi/eebus-go-cem/features"
	"github.com/DerAndereAndi/eebus-go/spine"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

// Internal EventHandler Interface for the CEM
func (e *EV) HandleEvent(payload spine.EventPayload) {
	switch payload.EventType {
	case spine.EventTypeEntityChange:
		switch payload.ChangeType {
		case spine.ElementChangeAdd:
			// EV connected

			if !e.checkEntityBeingEV(payload.Entity) {
				return
			}
			e.evConnected(payload.Entity)
		case spine.ElementChangeRemove:
			// EV disconnected
			if !e.checkEntityBeingEV(payload.Entity) {
				return
			}
			e.evDisconnected(payload.Entity)
		}
	case spine.EventTypeDataChange:
		if payload.ChangeType == spine.ElementChangeUpdate {
			switch payload.Data.(type) {
			case *model.DeviceConfigurationKeyValueDescriptionListDataType:
				// key value descriptions received, now get the data
				_, _ = features.RequestDeviceConfigurationKeyValueList(e.service, payload.Entity)

			case *model.DeviceDiagnosisStateDataType:
				// TODO: received diagnosis state

			}
		}
	}
}

// check if the provided entity is an EV
func (e *EV) checkEntityBeingEV(entity *spine.EntityRemoteImpl) bool {
	if entity == nil || entity.EntityType() != model.EntityTypeTypeEV {
		return false
	}
	return true
}

// an EV was disconnected, trigger required cleanup
func (e *EV) evDisconnected(entity *spine.EntityRemoteImpl) {
	fmt.Println("EV DISCONNECTED")

	// TODO: add error handling

}

// an EV was connected, trigger required communication
func (e *EV) evConnected(entity *spine.EntityRemoteImpl) {
	fmt.Println("EV CONNECTED")

	// TODO: add error handling

	// get ev configuration data
	if err := features.RequestDeviceConfiguration(e.service, entity); err != nil {
		return
	}

	// get manufacturer details
	if _, err := features.RequestManufacturerDetailsForEntity(e.service, entity); err != nil {
		return
	}

	// get device diagnosis state
	if _, err := features.RequestDiagnosisStateForEntity(e.service, entity); err != nil {
		return
	}

	// get electrical connection parameter
	if err := features.RequestElectricalConnection(e.service, entity); err != nil {
		return
	}

	// get measurement parameters
	if err := features.RequestMeasurement(e.service, entity); err != nil {
		return
	}

	// get identification
	if err := features.RequestIdentification(e.service, entity); err != nil {
		return
	}

}