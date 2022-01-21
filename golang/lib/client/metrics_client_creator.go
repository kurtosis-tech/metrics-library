package client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/do_nothing_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/snow_plow_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
)

func CreateDefaultMetricsClient(source source.Source, userId string, usserAcceptSendingMetrics bool) (MetricsClient, error) {

	metricsProvider := DoNoting

	if usserAcceptSendingMetrics{
		metricsProvider = SnowPlow
	}

	return CreateMetricsClient(source, userId, metricsProvider)
}

func CreateMetricsClient(source source.Source, userId string, metricsProvider MetricsClientProvider) (MetricsClient, error) {

	switch metricsProvider {
	case SnowPlow:
		metricsClient, err := snow_plow_client.NewSnowPlowClient(source, userId)
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred creating SnowPlow metrics client")
		}
		return metricsClient, nil
	case DoNoting:
		metricsClient := do_nothing_client.NewDoNothingClient()
		return metricsClient, nil
    default:
		return nil, stacktrace.NewError("Unrecognized metrics provider '%v'", metricsProvider)
	}
}

