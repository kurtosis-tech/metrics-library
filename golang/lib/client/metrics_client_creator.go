package client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/do_nothing_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/segment_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
	"gopkg.in/segmentio/analytics-go.v3"
)

const(
	defaultMetricsType = Segment
)

//The argument shouldFlushQueueOnEachEvent is used to imitate a sync request, it is not exactly the same because
//the event is enqueued but the queue is flushed suddenly so is pretty close to event traked in sync
//The argument callbackObject is an object that will be used by the client to notify the
// application when messages sends to the backend API succeeded or failed.
func CreateMetricsClient(source source.Source, sourceVersion string, userId string, didUserAcceptSendingMetrics bool, shouldFlushQueueOnEachEvent bool, callbackObject analytics.Callback) (MetricsClient, func() error, error) {

	metricsClientType := DoNothing

	if didUserAcceptSendingMetrics{
		metricsClientType = defaultMetricsType
	}

	switch metricsClientType {
	case Segment:
		metricsClient, err := segment_client.NewSegmentClient(source, sourceVersion, userId, shouldFlushQueueOnEachEvent, callbackObject)
		if err != nil {
			return nil, nil, stacktrace.Propagate(err, "An error occurred creating Segment metrics client")
		}
		return metricsClient, metricsClient.Close, nil
	case DoNothing:
		metricsClient := do_nothing_client.NewDoNothingClient()
		return metricsClient, metricsClient.Close, nil
	default:
		return nil, nil, stacktrace.NewError("Unrecognized metrics client type '%v'", metricsClientType)
	}
}
