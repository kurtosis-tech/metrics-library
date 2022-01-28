package event

import (
	"crypto/sha256"
	"fmt"
	"github.com/kurtosis-tech/stacktrace"
	"strings"
)

const (
	yesStr = "yes"
	noStr = "no"


	//We are following these naming conventions for event's data
	//https://segment.com/docs/getting-started/04-full-install/#event-naming-best-practices
	enclaveIDPropertyKey = "enclave_id"
	moduleIDPropertyKey = "module_id"
	containerImageNamePropertyKey = "container_image_name"
	containerImageVersionPropertyKey = "container_image_version"
	moduleParamsPropertyKey = "module_params"
	didUserAcceptSendingMetricsKey = "did_user_accept_sending_metrics"

	//Categories
	installCategory = "install"
	enclaveCategory = "enclave"
	moduleCategory = "module"

	//Actions
	consentAction = "consent"
	createAction = "create"
	stopAction = "stop"
	destroyAction = "destroy"
	loadAction = "load"
	unloadAction = "unload"
	executeAction = "execute"

	containerImageSeparatorCharacter = ":"
	validAmountOfColonsInContainerImage = 1

)

func NewShouldSendMetricsUserElectionEvent(didUserAcceptSendingMetrics bool) (*Event, error) {

	didUserAcceptSendingMetricsStr := fmt.Sprintf("%v", didUserAcceptSendingMetrics)

	properties := map[string]string{
		didUserAcceptSendingMetricsKey: didUserAcceptSendingMetricsStr,
	}

	event, err := newEvent(installCategory, consentAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new user accept sending metrics event")
	}

	return event, nil
}

func NewCreateEnclaveEvent(enclaveId string) (*Event, error) {
	hashedEnclaveId, err := chekIfNotEmptyStringAndGetHashedValue(enclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating and getting hashed enclave id")
	}

	properties := map[string]string{
		enclaveIDPropertyKey: hashedEnclaveId,
	}

	event, err := newEvent(enclaveCategory, createAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new create enclave event")
	}

	return event, nil
}

func NewStopEnclaveEvent(enclaveId string) (*Event, error) {
	hashedEnclaveId, err := chekIfNotEmptyStringAndGetHashedValue(enclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating and getting hashed enclave id")
	}

	properties := map[string]string{
		enclaveIDPropertyKey: hashedEnclaveId,
	}

	event, err := newEvent(enclaveCategory, stopAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new stop enclave event")
	}

	return event, nil
}

func NewDestroyEnclaveEvent(enclaveId string) (*Event, error) {
	hashedEnclaveId, err := chekIfNotEmptyStringAndGetHashedValue(enclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Enclave ID can not be an empty string")
	}

	properties := map[string]string{
		enclaveIDPropertyKey: hashedEnclaveId,
	}

	event, err := newEvent(enclaveCategory, destroyAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new destroy enclave event")
	}

	return event, nil
}

func NewLoadModuleEvent(moduleId, containerImage, serializedParams string) (*Event, error) {

	hashedModuleId, err := chekIfNotEmptyStringAndGetHashedValue(moduleId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Module ID can not be an empty string")
	}

	containerImageName, containerImageVersion, err := splitContainerImageIntoNameAndVersion(containerImage)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred splitting container image '%v' into image name and image version", containerImage)
	}

	hashedSerializedParams := hashString(serializedParams)

	properties := map[string]string{
		moduleIDPropertyKey: hashedModuleId,
		containerImageNamePropertyKey: containerImageName,
		containerImageVersionPropertyKey: containerImageVersion,
		moduleParamsPropertyKey: hashedSerializedParams,
	}

	event, err := newEvent(moduleCategory, loadAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new load module event")
	}
	return event, nil
}

func NewUnloadModuleEvent(moduleId string) (*Event, error) {

	hashedModuleId, err := chekIfNotEmptyStringAndGetHashedValue(moduleId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Module ID can not be an empty string")
	}

	properties := map[string]string{
		moduleIDPropertyKey: hashedModuleId,
	}

	event, err := newEvent(moduleCategory, unloadAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new unload module event")
	}
	return event, nil
}

func NewExecuteModuleEvent(moduleId, serializedParams string) (*Event, error) {

	hashedModuleId, err := chekIfNotEmptyStringAndGetHashedValue(moduleId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Module ID can not be an empty string")
	}

	hashedSerializedParams := hashString(serializedParams)

	properties := map[string]string{
		moduleIDPropertyKey: hashedModuleId,
		moduleParamsPropertyKey: hashedSerializedParams,
	}

	event, err := newEvent(moduleCategory, executeAction, properties)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new execute module event")
	}
	return event, nil
}

// ================================================================================================
//                                  Private Helper Functions
// ================================================================================================
func chekIfNotEmptyStringAndGetHashedValue(value string) (string, error)  {
	valueWithoutSpaces := strings.TrimSpace(value)

	if valueWithoutSpaces == "" {
		return "", stacktrace.NewError("Invalid value, it can not be an empty string")
	}

	hashedValue := hashString(value)

	return hashedValue, nil
}

func hashString(value string) string {
	hash := sha256.New()

	hash.Write([]byte(value))

	hashedByteSlice := hash.Sum(nil)

	hexValue := fmt.Sprintf("%x", hashedByteSlice)

	return hexValue
}

func splitContainerImageIntoNameAndVersion(containerImage string) (string, string, error) {
	amountOfColons := strings.Count(containerImage, containerImageSeparatorCharacter)
	if amountOfColons != validAmountOfColonsInContainerImage{
		return "", "", stacktrace.NewError("Invalid container image, it has '%v' colons and should has '%v'", amountOfColons, validAmountOfColonsInContainerImage)
	}

	containerImageSlice := strings.Split(containerImage, containerImageSeparatorCharacter)

	return containerImageSlice[0], containerImageSlice[1], nil
}
