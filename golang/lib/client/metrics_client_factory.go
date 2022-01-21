package client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/snow_plow_client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
)

const (
	defaultMetricsClientProvider = SnowPlow
)

func CreateDefaultMetricsClient(source source.Source, userId string) (MetricsClient, error) {
	return CreateMetricsClient(source, userId, defaultMetricsClientProvider)
}

func CreateMetricsClient(source source.Source, userId string, metricsProvider MetricsClientProvider) (MetricsClient, error) {

	switch metricsProvider {
	case SnowPlow:
		metricsClient, err := snow_plow_client.NewSnowPlowClient(source, userId)
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred creating SnowPlow metrics client")
		}
		return metricsClient, nil
    default:
		return nil, stacktrace.NewError("Unrecognized metrics provider '%v'", metricsProvider)
	}
}

