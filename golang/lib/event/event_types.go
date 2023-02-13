package event

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	//We are following these naming conventions for event's data
	//https://segment.com/docs/getting-started/04-full-install/#event-naming-best-practices
	enclaveIDPropertyKey           = "enclave_id"
	didUserAcceptSendingMetricsKey = "did_user_accept_sending_metrics"
	packageIdKey                   = "package_id"
	isRemotePackageKey             = "is_remote_package"
	isDryRunKey                    = "is_dry_run"
	isScriptKey                    = "is_script"

	//Categories
	installCategory = "install"
	enclaveCategory = "enclave"
	// the Kurtosis category is for commands at the root level of the cli
	// we went this way cause this is in pattern with other categories above
	// any further root level commands should use this category
	kurtosisCategory = "kurtosis"

	//Actions
	consentAction = "consent"
	createAction  = "create"
	stopAction    = "stop"
	destroyAction = "destroy"
	runAction     = "run"

	containerImageSeparatorCharacter    = ":"
	validAmountOfColonsInContainerImage = 1
	dockerDefaultImageTag               = "latest"
)

// WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING
// NO EVENTS SHOULD RETURN AN ERROR! Instead, each event should *always* return an event (even if the value is garbage)
// This is becasue if we return an error, the error will propagate which means that we don't even send the event
//  at all, which means that we'll silently drop data, which means we won't even realize that something is wrong!
// If we send the event with garbage data, it means that we at least get the chance to notice it in our product analytics
//  dashboards.
// WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING

func NewShouldSendMetricsUserElectionEvent(didUserAcceptSendingMetrics bool) *Event {
	didUserAcceptSendingMetricsStr := fmt.Sprintf("%v", didUserAcceptSendingMetrics)
	properties := map[string]string{
		didUserAcceptSendingMetricsKey: didUserAcceptSendingMetricsStr,
	}
	event := newEvent(installCategory, consentAction, properties)
	return event
}

func NewCreateEnclaveEvent(enclaveId string) *Event {
	hashedEnclaveId := hashString(strings.TrimSpace(enclaveId))
	properties := map[string]string{
		enclaveIDPropertyKey: hashedEnclaveId,
	}
	event := newEvent(enclaveCategory, createAction, properties)
	return event
}

func NewStopEnclaveEvent(enclaveId string) *Event {
	hashedEnclaveId := hashString(strings.TrimSpace(enclaveId))
	properties := map[string]string{
		enclaveIDPropertyKey: hashedEnclaveId,
	}
	event := newEvent(enclaveCategory, stopAction, properties)
	return event
}

func NewDestroyEnclaveEvent(enclaveId string) *Event {
	hashedEnclaveId := hashString(strings.TrimSpace(enclaveId))
	properties := map[string]string{
		enclaveIDPropertyKey: hashedEnclaveId,
	}
	event := newEvent(enclaveCategory, destroyAction, properties)
	return event
}

func NewKurtosisRunEvent(packageId string, isRemote bool, isDryRun bool, isScript bool) *Event {
	isRemotePackageStr := fmt.Sprintf("%v", isRemote)
	isDryRunStr := fmt.Sprintf("%v", isDryRun)
	isScriptStr := fmt.Sprintf("%v", isScript)

	properties := map[string]string{
		packageIdKey:       packageId,
		isRemotePackageKey: isRemotePackageStr,
		isDryRunKey:        isDryRunStr,
		isScriptKey:        isScriptStr,
	}

	event := newEvent(kurtosisCategory, runAction, properties)
	return event
}

// ================================================================================================
//
//	Private Helper Functions
//
// ================================================================================================
func hashString(value string) string {
	hash := sha256.New()

	hash.Write([]byte(value))

	hashedByteSlice := hash.Sum(nil)

	hexValue := fmt.Sprintf("%x", hashedByteSlice)

	return hexValue
}
