package client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/do_nothing_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/segment_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
)

const(
	defaultMetricsType = Segment
)

func CreateMetricsClient(source source.Source, sourceVersion string, userId string, didUserAcceptSendingMetrics bool, shouldFlushQueueOnEachEvent bool) (MetricsClient, error) {

	metricsClientType := DoNothing

	if didUserAcceptSendingMetrics{
		metricsClientType = defaultMetricsType
	}

	switch metricsClientType {
	case Segment:
		metricsClient, err := segment_client.NewSegmentClient(source, sourceVersion, userId, shouldFlushQueueOnEachEvent)
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred creating Segment metrics client")
		}
		return metricsClient, nil
	case DoNothing:
		metricsClient := do_nothing_client.NewDoNothingClient()
		return metricsClient, nil
	default:
		return nil, stacktrace.NewError("Unrecognized metrics client type '%v'", metricsClientType)
	}
}
