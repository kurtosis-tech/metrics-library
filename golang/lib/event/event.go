package event

import (
	"github.com/kurtosis-tech/stacktrace"
	"strings"
)

type Event struct {
	//Category of event (e.g. enclave, module)
	category string

	//Action performed/event name (e.g. create, load)
	action string

	//A property Key associated with the object of the action (e.g. enclave ID, module name)
	propertyKey string

	//The property value
	propertyValue string
}

func newEvent(category, action, propertyKey, propertyValue string) (*Event, error) {

	event := &Event{category: category, action: action, propertyKey: propertyKey, propertyValue: propertyValue}

	if err := event.IsValid(); err != nil {
		return nil, stacktrace.Propagate(err, "Invalid event '%+v'", event)
	}

	return event, nil
}

func (event *Event) GetCategory() string {
	return event.category
}

func (event *Event) GetAction() string {
	return event.action
}

func (event *Event) GetName() string {
	return strings.Join([]string{event.category,event.action}, "-")
}

func (event *Event) GetPropertyKey() string {
	return event.propertyKey
}

func (event *Event) GetPropertyValue() string {
	return event.propertyValue
}

//IsValid return nil if the event is valid
//Category an action are mandatory
func (event *Event) IsValid() error {
	category := strings.TrimSpace(event.category)

	if category != "" {
		return stacktrace.NewError("Event's category can not be empty string")
	}

	action := strings.TrimSpace(event.action)

	if action != "" {
		return stacktrace.NewError("Event's action can not be empty string")
	}

	return nil
}
