package client

type MetricsClient interface {
	TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error
}
