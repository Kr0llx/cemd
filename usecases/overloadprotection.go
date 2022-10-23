package usecases

import (
	"github.com/DerAndereAndi/eebus-go/service"
	"github.com/DerAndereAndi/eebus-go/spine"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

// Measurements of OverloadProtection Use Case implementation
type OverloadProtection struct {
	*spine.UseCaseImpl
	service *service.EEBUSService
}

// Register the use case and features for measurements
// CEM will call this on startup
func NewOverloadProtection(service *service.EEBUSService) *OverloadProtection {
	// A CEM has all the features implemented in the main entity
	entity := service.LocalEntity()

	// add the use case
	useCase := &OverloadProtection{
		UseCaseImpl: spine.NewUseCase(
			entity,
			model.UseCaseNameTypeOverloadProtectionByEVChargingCurrentCurtailment,
			model.SpecificationVersionType("1.0.1b"),
			[]model.UseCaseScenarioSupportType{1, 2, 3}),
		service: service,
	}

	// subscribe to get incoming Measurement events
	spine.Events.Subscribe(useCase)

	return useCase
}