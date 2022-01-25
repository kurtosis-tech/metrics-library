package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveId string) error
	TrackStopEnclave(enclaveId string) error
	TrackDestroyEnclave(enclaveId string) error
	TrackCleanEnclave(shouldCleanAll bool) error
	TrackLoadModule(moduleId string) error
	TrackExecuteModule(moduleId string) error
	TrackUnloadModule(moduleId string) error
}
