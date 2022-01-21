package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveId string) error
	TrackStopEnclave() error
	TrackDestroyEnclave() error
	TrackCleanEnclave() error
}
