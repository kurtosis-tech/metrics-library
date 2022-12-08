package client

type MetricsClient interface {
	TrackShouldSendMetricsUserElection(didUserAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveId string) error
	TrackStopEnclave(enclaveId string) error
	TrackDestroyEnclave(enclaveId string) error
	TrackLoadModule(moduleId, containerImage, serializedParams string) error
	TrackExecuteModule(moduleId, serializedParams string) error
	TrackUnloadModule(moduleId string) error
	TrackRunStarlarkPackage(isRemote bool, packageId, serializedArgs string, isDryRun bool) error
	TrackRunStarlarkScript(serializedScript string, serializedArgs string, isDryRun bool) error
	close() (err error)
}
