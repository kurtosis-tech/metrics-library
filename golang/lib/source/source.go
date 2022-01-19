package source

import "github.com/kurtosis-tech/stacktrace"

const (
	KurtosisCLISource    Source = "kurtosis-cli"
	KurtosisEngineSource Source = "kurtosis-engine"
	KurtosisAPISource    Source = "kurtosis-api"
)

var allValidSources = map[Source]bool{
	KurtosisCLISource:    true,
	KurtosisEngineSource: true,
	KurtosisAPISource:    true,
}

type Source string

func (source Source) IsValid() error {
	if _, found := allValidSources[source]; !found {
		return stacktrace.NewError("The source '%v' is not valid. Valid sources: %+v", source, allValidSources)
	}
	return nil
}
