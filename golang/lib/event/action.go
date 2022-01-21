package event

import "github.com/kurtosis-tech/stacktrace"

const (
	CreateAction  Action = "create"
	DestroyAction Action = "destroy"
	StopAction    Action = "stop"
	LoadAction    Action = "load"
	UnloadAction  Action = "unload"
	ConsentAction Action = "consent"
)

var allValidActions = map[Action]bool{
	CreateAction:  true,
	DestroyAction: true,
	StopAction:    true,
	LoadAction:    true,
	UnloadAction:  true,
}

var allValidEnclaveActions = map[Action]bool{
	CreateAction:  true,
	DestroyAction: true,
	StopAction:    true,
}

var allValidModuleActions = map[Action]bool{
	LoadAction:   true,
	UnloadAction: true,
}

var allValidInstallActions = map[Action]bool{
	ConsentAction: true,
}

var allValidActionsByCategory = map[Category]map[Action]bool{
	EnclaveCategory: allValidEnclaveActions,
	ModuleCategory:  allValidModuleActions,
	InstallCategory: allValidInstallActions,
}

type Action string

func (action Action) IsValid() error {
	if _, found := allValidActions[action]; !found {
		return stacktrace.NewError("The action '%v' is not valid. Valid actions: %+v", action, allValidActions)
	}
	return nil
}

func (action Action) IsValidForCategory(category Category) error {
	validActionsForCategory, found := allValidActionsByCategory[category]
	if !found {
		return stacktrace.NewError("The category '%v' haven't valid actions")
	}

	if _, found = validActionsForCategory[action]; !found {
		return stacktrace.NewError("The action '%v' is not valid for category %v. Valid actions: %+v", action, category, validActionsForCategory)
	}

	return nil
}
