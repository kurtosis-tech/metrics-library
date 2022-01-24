package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveId string) error
	TrackStopEnclave(enclaveId string) error
	TrackDestroyEnclave(enclaveId string) error
	TrackCleanEnclave(shouldCleanAll bool) error
	//TODO kurtosis'module tracking
}
