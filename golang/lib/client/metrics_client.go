package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
	//This method must allow us to disable tracking any time
	//The implementation should guarantee that no more metrics will be sent
	DisableTracking()
}
