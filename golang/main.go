package main

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/client"
	"github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {

	metricsClient, err := client.CreateDefaultMetricsClient(source.KurtosisEngineSource, "1.2.3", "Leo-Testing-1-2-3", true)
	if err != nil {
		logrus.Infof("An error occurred creating SnowPlow metrics client, error \n%v", err)
	}

	if err := metricsClient.TrackLoadModule("my-module-id"); err != nil {
		logrus.Infof("An error occurred tracking load module")
	}

	time.Sleep(5 * time.Minute)
}
