package client

import "github.com/kurtosis-tech/metrics-library/golang/lib/event"

type MetricsClient interface {
	Track(event *event.Event) error
	//This method must allow us to disable tracking any time
	//The implementation should guarantee that no more metrics will be sent
	DisableTracking()
}
