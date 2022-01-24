package segment_client

import (
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

	//We should uncomment this code if we want to create an event to identify the user
	//every time the client is created, it will be add a new row in SF "Identifies" table
	/*
	if err := client.Enqueue(analytics.Identify{
		UserId: userId,
	}); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred enqueuing a new identify event in Segment client's queue")
	}*/

	return &SegmentClient{client: client, source: source, userID: userId}, nil
}


func (segment *SegmentClient) TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error {

	newEvent, err := event.NewUserAcceptSendingMetricsEvent(userAcceptSendingMetrics)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new user accept sending metrics event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking user accept sending metrics event")
	}

	return nil
}

func (segment *SegmentClient) TrackCreateEnclave(enclaveId string) error {

	newEvent, err := event.NewCreateEnclaveEvent(enclaveId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new create enclave event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking create enclave event")
	}

	return nil
}

func (segment *SegmentClient) TrackStopEnclave(enclaveId string) error {
	newEvent, err := event.NewStopEnclaveEvent(enclaveId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new stop enclave event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking stop enclave event")
	}

	return nil
}

func (segment *SegmentClient) TrackDestroyEnclave(enclaveId string) error {
	newEvent, err := event.NewDestroyEnclaveEvent(enclaveId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new destroy enclave event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking destroy enclave event")
	}
	return nil
}

func (segment *SegmentClient) TrackCleanEnclave(shouldCleanAll bool) error {
	newEvent, err := event.NewCleanEnclaveEvent(shouldCleanAll)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new clean enclave event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking clean enclave event")
	}

	return nil
}

func (segment *SegmentClient) track(event *event.Event) error {
	if err := segment.client.Enqueue(analytics.Track{
		Event:  event.GetName(),
		UserId: segment.userID,
		Properties: analytics.NewProperties().
			Set(event.GetPropertyKey(), event.GetPropertyValue()),
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}
	return nil
}
