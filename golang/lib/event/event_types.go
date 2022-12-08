package event

import (
	"crypto/sha256"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	//We are following these naming conventions for event's data
	//https://segment.com/docs/getting-started/04-full-install/#event-naming-best-practices
	enclaveIDPropertyKey               = "enclave_id"
	moduleIDPropertyKey                = "module_id"
	containerRawImageStringPropertyKey = "container_raw_image_string"
	containerImageNamePropertyKey      = "container_image_name"
	containerImageVersionPropertyKey   = "container_image_version"
	moduleParamsPropertyKey            = "module_params"
	didUserAcceptSendingMetricsKey     = "did_user_accept_sending_metrics"
	starlarkArgsKey                    = "starlark_args"
	packageIdKey                       = "package_id"
	isRemotePackageKey                 = "is_remote_package"
	starlarkSerializedScriptKey        = "serialized_script"
	isDryRunKey                        = "is_dry_run"

	//Categories
	installCategory         = "install"
	enclaveCategory         = "enclave"
	moduleCategory          = "module"
	starlarkPackageCategory = "package"
	starlarkScriptCategory  = "script"

	//Actions
	consentAction = "consent"
	createAction  = "create"
	stopAction    = "stop"
	destroyAction = "destroy"
	loadAction    = "load"
	unloadAction  = "unload"
	executeAction = "execute"
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

func NewLoadModuleEvent(moduleId, containerImage, serializedParams string) *Event {
	hashedModuleId := hashString(strings.TrimSpace(moduleId))

	containerImageWithoutSpaces := strings.TrimSpace(containerImage)
	containerImageOrgAndRepo, containerImageTag := bestEffortSplitContainerImageIntoOrgRepoAndVersion(containerImageWithoutSpaces)

	hashedSerializedParams := hashString(serializedParams)
	properties := map[string]string{
		moduleIDPropertyKey:                hashedModuleId,
		containerRawImageStringPropertyKey: containerImage,
		containerImageNamePropertyKey:      containerImageOrgAndRepo,
		containerImageVersionPropertyKey:   containerImageTag,
		moduleParamsPropertyKey:            hashedSerializedParams,
	}

	event := newEvent(moduleCategory, loadAction, properties)
	return event
}

func NewUnloadModuleEvent(moduleId string) *Event {
	hashedModuleId := hashString(strings.TrimSpace(moduleId))
	properties := map[string]string{
		moduleIDPropertyKey: hashedModuleId,
	}
	event := newEvent(moduleCategory, unloadAction, properties)
	return event
}

func NewExecuteModuleEvent(moduleId, serializedParams string) *Event {
	hashedModuleId := hashString(strings.TrimSpace(moduleId))

	hashedSerializedParams := hashString(serializedParams)

	properties := map[string]string{
		moduleIDPropertyKey:     hashedModuleId,
		moduleParamsPropertyKey: hashedSerializedParams,
	}

	event := newEvent(moduleCategory, executeAction, properties)
	return event
}

func NewRunStarlarkPackage(packageId string, serializedArgs string, isRemote bool, isDryRun bool) *Event {
	hashedPackageId := hashString(strings.TrimSpace(packageId))
	hashedSerializedArgs := hashString(strings.TrimSpace(serializedArgs))
	isRemotePackageStr := fmt.Sprintf("%v", isRemote)
	isDryRunStr := fmt.Sprintf("%v", isDryRun)

	properties := map[string]string{
		starlarkArgsKey:    hashedSerializedArgs,
		packageIdKey:       hashedPackageId,
		isRemotePackageKey: isRemotePackageStr,
		isDryRunStr:        isDryRunStr,
	}

	event := newEvent(starlarkPackageCategory, runAction, properties)
	return event
}

func NewRunStarlarkScript(serializedScript string, serializedArgs string, isDryRun bool) *Event {
	hashedSerializedScript := hashString(strings.TrimSpace(serializedScript))
	hashedSerializedArgs := hashString(strings.TrimSpace(serializedArgs))
	isDryRunStr := fmt.Sprintf("%v", isDryRun)

	properties := map[string]string{
		starlarkArgsKey:             hashedSerializedArgs,
		starlarkSerializedScriptKey: hashedSerializedScript,
		isDryRunStr:                 isDryRunStr,
	}
	event := newEvent(starlarkScriptCategory, runAction, properties)
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

// Makes a best-effort attempt to split the container image spec into org-and-repo and tag, returning
//
//	emptystrings if it's not successful
func bestEffortSplitContainerImageIntoOrgRepoAndVersion(containerImage string) (resultOrgAndRepo string, resultTag string) {
	amountOfColons := strings.Count(containerImage, containerImageSeparatorCharacter)
	if amountOfColons == 0 {
		containerImage = containerImage + containerImageSeparatorCharacter + dockerDefaultImageTag
	}
	if amountOfColons > validAmountOfColonsInContainerImage {
		logrus.Debugf("Invalid container image '%v', it has '%v' colons and should have '%v'", containerImage, amountOfColons, validAmountOfColonsInContainerImage)
		return "", ""
	}

	containerImageSlice := strings.Split(containerImage, containerImageSeparatorCharacter)

	return containerImageSlice[0], containerImageSlice[1]
}
