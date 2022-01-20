package event

import "github.com/kurtosis-tech/stacktrace"

type Event struct {
	//Category of event (e.g. enclave, module)
	category Category

	//Action performed/event name (e.g. create, load)
	action Action

	//The object of the action (e.g. enclave ID, module name)
	label string

	//A property associated with the object of the action (e.g. partitioning-enabled)
	property string

	//A value associated with the event/action (in most cases it will be ignored because
	//we are going to track individual events, such as: create enclave, only creates 1 enclave
	//load module, only loads 1 module. But could be the case that we want so send a value)
	value float64
}

func (event *Event) GetCategory() Category {
	return event.category
}

func (event *Event) GetCategoryString() string {
	return string(event.category)
}

func (event *Event) GetAction() Action {
	return event.action
}

func (event *Event) GetActionString() string {
	return string(event.action)
}

func (event *Event) GetLabel() string {
	return event.label
}

func (event *Event) GetProperty() string {
	return event.property
}

func (event *Event) GetValue() float64 {
	return event.value
}

//IsValid return nil if the event is valid
func (event *Event) IsValid() error {
	if err := event.category.IsValid(); err != nil {
		return stacktrace.Propagate(err, "Invalid event's category")
	}

	if err := event.action.IsValidForCategory(event.category); err != nil {
		return stacktrace.Propagate(err, "Invalid event's action")
	}

	return nil
}


// ====================================================================================================
//                                      Builder
// ====================================================================================================
type EventBuilder struct {
	category Category
	action   Action
	label    string
	property string
	value    float64
}

func NewEventBuilder(category Category, action Action) *EventBuilder {
	return &EventBuilder{
		category: category,
		action:   action,
		label:    "",
		property: "",
		value:    0,
	}
}

func (builder *EventBuilder) WithLabel(label string) *EventBuilder {
	builder.label = label
	return builder
}

func (builder *EventBuilder) WithProperty(property string) *EventBuilder {
	builder.property = property
	return builder
}

func (builder *EventBuilder) WithValue(value float64) *EventBuilder {
	builder.value = value
	return builder
}

func (builder *EventBuilder) Build() (*Event, error) {
	event := &Event{
		category: builder.category,
		action:   builder.action,
		label:    builder.label,
		property: builder.property,
		value:    builder.value,
	}

	if err := event.IsValid(); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating the event '%+v'", event)
	}

	return event, nil
}
