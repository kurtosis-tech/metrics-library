package segment_client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/event"
	metrics_source "github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/segmentio/backo-go"
	"gopkg.in/segmentio/analytics-go.v3"
	"time"
)

const (
	//Key generated in my lporoli trial account
	accountWriteKey = "WbfsEYlBdRyaML5adTucEzqBkpQsz4p7"

	shouldTrackIdentifyUserEventWhenClientIsCreated = false
)

type SegmentClient struct {
	client analytics.Client
	analyticsContext *analytics.Context
	userID string
}

func NewSegmentClient(source metrics_source.Source, sourceVersion string, userId string) (*SegmentClient, error) {
	if err := source.IsValid(); err != nil {
		return nil, stacktrace.Propagate(err, "Invalid source")
	}

	config := analytics.Config{
		Interval: 10 * time.Minute,
		//NOTE: Segment client has a max attempt = 10, so this retry strategy
		//allow us to execute the first attempt in 5 seconds and the last attend in 24 hours
		//which is useful if a user is executing the metrics without internet connection for several hours
		RetryAfter: func(attempt int) time.Duration {
			retryBacko := backo.NewBacko(time.Second*5, 3, 0, time.Hour*24)
			return retryBacko.Duration(attempt)
		},
	}

	client, err := analytics.NewWithConfig(accountWriteKey, config)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating new Segment client with config '%+v'", config)
	}

	analyticsContext := newAnalyticsContext(source, sourceVersion)

	//We could activate this functionality if we want to track an event to identify the user
	//every time the client is created, it will be adding a new row in SF "Identifies" table
	if shouldTrackIdentifyUserEventWhenClientIsCreated {
		if err := client.Enqueue(analytics.Identify{
			UserId: userId,
			Context: analyticsContext,
		}); err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred enqueuing a new identify event in Segment client's queue")
		}
	}

	return &SegmentClient{client: client, analyticsContext: analyticsContext, userID: userId}, nil
}


func (segment *SegmentClient) TrackUserAcceptSendingMetrics(didUserAcceptSendingMetrics bool) error {

	newEvent, err := event.NewUserAcceptSendingMetricsEvent(didUserAcceptSendingMetrics)
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

func (segment *SegmentClient) TrackLoadModule(moduleId string) error {
	newEvent, err := event.NewLoadModuleEvent(moduleId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new load module event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking load module event")
	}

	return nil
}

func (segment *SegmentClient) TrackUnloadModule(moduleId string) error {
	newEvent, err := event.NewUnloadModuleEvent(moduleId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new unload module event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking unload module event")
	}

	return nil
}

func (segment *SegmentClient) TrackExecuteModule(moduleId string) error {
	newEvent, err := event.NewExecuteModuleEvent(moduleId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new execute module event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking execute module event")
	}

	return nil
}

// ====================================================================================================
// 									   Private helper methods
// ====================================================================================================
func (segment *SegmentClient) track(event *event.Event) error {
	if err := segment.client.Enqueue(analytics.Track{
		Event:  event.GetName(),
		UserId: segment.userID,
		Context: segment.analyticsContext,
		Properties: analytics.NewProperties().
			Set(event.GetPropertyKey(), event.GetPropertyValue()),
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}
	return nil
}

func newAnalyticsContext(source metrics_source.Source, sourceVersion string) *analytics.Context {
	appInfo := analytics.AppInfo{
		Name: string(source),
		Version: sourceVersion,
	}

	analyticsContext := &analytics.Context{App: appInfo}

	return analyticsContext
}
