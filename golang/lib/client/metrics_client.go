package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveId string) error
	TrackStopEnclave() error  //TODO check if we can send the enclaveID
	TrackDestroyEnclave() error //TODO check if we can send the enclaveID
	TrackCleanEnclave() error //TODO check if we can send the enclaveID
	//TODO kurtosis'module tracking
}
