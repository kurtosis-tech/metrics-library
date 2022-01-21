package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveName string) error
	TrackStopEnclave() error
	TrackDestroyEnclave() error
}
