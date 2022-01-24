package segment_client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/client/common"
	"github.com/kurtosis-tech/metrics-library/golang/lib/event"
	metrics_source "github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
	"gopkg.in/segmentio/analytics-go.v3"
)

const (
	//Key generated in my lporoli trial account
	accountWriteKey = "WbfsEYlBdRyaML5adTucEzqBkpQsz4p7"

	yesStr = "yes"
	noStr = "no"
)

type SegmentClient struct {
	client analytics.Client
	source metrics_source.Source
	userID string
}

func NewSegmentClient(source metrics_source.Source, userId string) (*SegmentClient, error) {

	client := analytics.New(accountWriteKey)

	//this call should be executed only the first time and also if some traits change
	/*
	if err := client.Enqueue(analytics.Identify{
		UserId: userId,
	}); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred enqueuing a new identify event in Segment client's queue")
	}*/

	return &SegmentClient{client: client, source: source, userID: "Leo-test"}, nil
}


func (segment *SegmentClient) TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error {

	var metricsLabel string
	if userAcceptSendingMetrics{
		metricsLabel = yesStr
	} else {
		metricsLabel = noStr
	}

	if err := segment.client.Enqueue(analytics.Track{
		Event:  string(event.InstallCategory) + "-" + string(event.ConsentAction),
		UserId: segment.userID,
		Properties: analytics.NewProperties().
			Set("user-accept-sending-metrics", metricsLabel),
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}

	return nil
}

func (segment *SegmentClient) TrackCreateEnclave(enclaveId string) error {

	hashedEnclaveId := common.HashString(enclaveId)

	if err := segment.client.Enqueue(analytics.Track{
		Event:  string(event.EnclaveCategory) + "-" + string(event.CreateAction),
		UserId: segment.userID,
		Properties: analytics.NewProperties().
			Set("enclave-id", hashedEnclaveId),
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}

	return nil
}

func (segment *SegmentClient) TrackStopEnclave() error {
	if err := segment.client.Enqueue(analytics.Track{
		Event:  string(event.EnclaveCategory) + "-" + string(event.StopAction),
		UserId: segment.userID,
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}
	return nil
}

func (segment *SegmentClient) TrackDestroyEnclave() error {
	if err := segment.client.Enqueue(analytics.Track{
		Event:  string(event.EnclaveCategory) + "-" + string(event.DestroyAction),
		UserId: segment.userID,
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}
	return nil
}

func (segment *SegmentClient) TrackCleanEnclave() error {
	if err := segment.client.Enqueue(analytics.Track{
		Event:  string(event.EnclaveCategory) + "-" + string(event.CleanAction),
		UserId: segment.userID,
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}
	return nil
}
