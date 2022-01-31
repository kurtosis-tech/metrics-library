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

	//Properties' keys and values associated with the object of the action (e.g. enclave ID, module name)
	properties map[string]string

}

func newEvent(category string, action string, properties map[string]string) (*Event, error) {

	categoryWithoutSpaces := strings.TrimSpace(category)

	if categoryWithoutSpaces == "" {
		return nil, stacktrace.NewError("Event's category can not be empty string")
	}

	actionWithoutSpaces := strings.TrimSpace(action)

	if actionWithoutSpaces == "" {
		return nil, stacktrace.NewError("Event's action can not be empty string")
	}

	for propertyKey := range properties {
		propertyKeyWithoutSpaces := strings.TrimSpace(propertyKey)
		if propertyKeyWithoutSpaces == "" {
			return nil, stacktrace.NewError("Propertie's key in an event can not be empty string")
		}
	}

	event := &Event{category: categoryWithoutSpaces, action: actionWithoutSpaces, properties: properties}

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

func (event *Event) GetProperties() map[string]string {
	return event.properties
}

