package client

import "github.com/kurtosis-tech/metrics-library/lib/event"

type MetricsClient interface {
	Track(event *event.Event) error
}
