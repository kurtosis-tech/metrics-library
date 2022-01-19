package event

import (
	"github.com/kurtosis-tech/stacktrace"
)

const (
	InstallCategory Category = "install"
	EnclaveCategory Category = "enclave"
	ModuleCategory  Category = "module"
)

var allValidCategories = map[Category]bool{
	EnclaveCategory: true,
	ModuleCategory:  true,
}

type Category string

func (category Category) IsValid() error {
	if _, found := allValidCategories[category]; !found {
		return stacktrace.NewError("The category '%v' is not valid. Valid categories: %+v", category, allValidCategories)
	}
	return nil
}
