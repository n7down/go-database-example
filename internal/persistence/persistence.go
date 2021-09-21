//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateInteraction(interaction Interaction) int64
	GetInteraction(id string) (recordNotFound bool, interaction Interaction)
	GetAllInteractions(limit, offset int32) (interaction []Interaction, err error)
	GetAllInteractionsByPagination(limit, offset int32) (GetAllInteractionsPagination, error)
	GetInteractionDetails(id string) ([]InteractionDetails, error)
	UpdateInteraction(interaction Interaction) (recordNotFound bool, err error)
	DeleteInteraction(interaction Interaction) (recordNotFound bool, err error)

	CreateKeypadCondition(keypadCondition KeypadCondition) int64
	GetKeypadCondition(id string) (recordNotFound bool, keypadCondition KeypadCondition)
	GetKeypadConditionByMac(mac string) (recordNotFound bool, keypadCondition KeypadCondition)
	GetKeypadConditionByMacAndButtonID(mac string, buttonID int) (recordNotFound bool, keypadCondition KeypadCondition)
	UpdateKeypadCondition(keypadCondition KeypadCondition) (recordNotFound bool, err error)
	DeleteKeypadCondition(keypadCondition KeypadCondition) (recordNotFound bool, err error)

	CreateLampEvent(lampEvent LampEvent) int64
	GetLampEvent(id string) (recordNotFound bool, lampEvent LampEvent)
	GetLampEventsByKeypadConditionID(id string) ([]LampEvent, error)
	GetLampEventsByKeypadMacAndButtonID(mac string, buttonID int) ([]LampEvent, error)
	UpdateLampEvent(lampEvent LampEvent) (recordNotFound bool, err error)
	DeleteLampEvent(lampEvent LampEvent) (recordNotFound bool, err error)

	CreateLampOnEvent(lampOnEvent LampOnEvent) int64
	GetLampOnEvent(id string) (recordNotFound bool, lampOnEvent LampOnEvent)
	UpdateLampOnEvent(lampOnEvent LampOnEvent) (recordNotFound bool, err error)
	DeleteLampOnEvent(lampOnEvent LampOnEvent) (recordNotFound bool, err error)

	CreateLampOffEvent(lampOffEvent LampOffEvent) int64
	GetLampOffEvent(id string) (recordNotFound bool, lampOffEvent LampOffEvent)
	UpdateLampOffEvent(lampOffEvent LampOffEvent) (recordNotFound bool, err error)
	DeleteLampOffEvent(lampOffEvent LampOffEvent) (recordNotFound bool, err error)

	CreateLampToggleEvent(lampToggleEvent LampToggleEvent) int64
	GetLampToggleEvent(id string) (recordNotFound bool, lampToggleEvent LampToggleEvent)
	UpdateLampToggleEvent(lampToggleEvent LampToggleEvent) (recordNotFound bool, err error)
	DeleteLampToggleEvent(lampToggleEvent LampToggleEvent) (recordNotFound bool, err error)

	CreateLampBrightnessEvent(lampBrightnessEvent LampBrightnessEvent) int64
	GetLampBrightnessEvent(id string) (recordNotFound bool, lampBrightnessEvent LampBrightnessEvent)
	UpdateLampBrightnessEvent(lampBrightnessEvent LampBrightnessEvent) (recordNotFound bool, err error)
	DeleteLampBrightnessEvent(lampBrightnessEvent LampBrightnessEvent) (recordNotFound bool, err error)

	CreateLampAutoBrightnessOnEvent(lampAutoBrightnessOnEvent LampAutoBrightnessOnEvent) int64
	GetLampAutoBrightnessOnEvent(id string) (recordNotFound bool, lampAutoBrightnessOnEvent LampAutoBrightnessOnEvent)
	UpdateLampAutoBrightnessOnEvent(lampAutoBrightnessOnEvent LampAutoBrightnessOnEvent) (recordNotFound bool, err error)
	DeleteLampAutoBrightnessOnEvent(lampAutoBrightnessOnEvent LampAutoBrightnessOnEvent) (recordNotFound bool, err error)

	CreateLampAutoBrightnessOffEvent(lampAutoBrightnessOffEvent LampAutoBrightnessOffEvent) int64
	GetLampAutoBrightnessOffEvent(id string) (recordNotFound bool, lampAutoBrightnessOffEvent LampAutoBrightnessOffEvent)
	UpdateLampAutoBrightnessOffEvent(lampAutoBrightnessOffEvent LampAutoBrightnessOffEvent) (recordNotFound bool, err error)
	DeleteLampAutoBrightnessOffEvent(lampAutoBrightnessOffEvent LampAutoBrightnessOffEvent) (recordNotFound bool, err error)

	CreateLampAutoBrightnessToggleEvent(lampAutoBrightnessToggleEvent LampAutoBrightnessToggleEvent) int64
	GetLampAutoBrightnessToggleEvent(id string) (recordNotFound bool, lampAutoBrightnessToggleEvent LampAutoBrightnessToggleEvent)
	UpdateLampAutoBrightnessToggleEvent(lampAutoBrightnessToggleEvent LampAutoBrightnessToggleEvent) (recordNotFound bool, err error)
	DeleteLampAutoBrightnessToggleEvent(lampAutoBrightnessToggleEvent LampAutoBrightnessToggleEvent) (recordNotFound bool, err error)

	CreateLampColorEvent(lampColorEvent LampColorEvent) int64
	GetLampColorEvent(id string) (recordNotFound bool, lampColorEvent LampColorEvent)
	UpdateLampColorEvent(lampColorEvent LampColorEvent) (recordNotFound bool, err error)
	DeleteLampColorEvent(lampColorEvent LampColorEvent) (recordNotFound bool, err error)

	CreateLampPulseEvent(lampPulseEvent LampPulseEvent) int64
	GetLampPulseEvent(id string) (recordNotFound bool, lampPulseEvent LampPulseEvent)
	UpdateLampPulseEvent(lampPulseEvent LampPulseEvent) (recordNotFound bool, err error)
	DeleteLampPulseEvent(lampPulseEvent LampPulseEvent) (recordNotFound bool, err error)

	CreateConditionsToEvents(conditionsToEvents ConditionsToEvents) int64
	GetConditionsToEvents(id string) (recordNotFound bool, conditionsToEvents ConditionsToEvents)
	UpdateConditionsToEvents(conditionsToEvents ConditionsToEvents) (recordNotFound bool, err error)
	DeleteConditionsToEvents(conditionsToEvents ConditionsToEvents) (recordNotFound bool, err error)

	CreateKeypadConditionToLampEvent(conditionsToEvents KeypadConditionsToLampEvents) int64
	GetKeypadConditionToLampEvent(id string) (recordNotFound bool, conditionsToEvents KeypadConditionsToLampEvents)
	UpdateKeypadConditionToLampEvent(conditionsToEvents KeypadConditionsToLampEvents) (recordNotFound bool, err error)
	DeleteKeypadConditionToLampEvent(conditionsToEvents KeypadConditionsToLampEvents) (recordNotFound bool, err error)
}
