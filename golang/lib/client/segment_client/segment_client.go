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

	segmentClientInterval = 10 * time.Minute

	retryBackoBaseDuration = time.Second*5
	retryBackoFactor = 3
	retryBackoJitter = 0
	retryBackoCap = time.Hour*24

	leastBatchSizeValue = 1
)

type SegmentClient struct {
	client analytics.Client
	analyticsContext *analytics.Context
	userID string
}

//TODO add a comment related to the shouldFlushQueueOnEachEvent argument
func NewSegmentClient(source metrics_source.Source, sourceVersion string, userId string, shouldFlushQueueOnEachEvent bool) (*SegmentClient, error) {

	config := analytics.Config{
		//The flushing interval of the client
		Interval: segmentClientInterval,
		//NOTE: Segment client has a max attempt = 10, so this retry strategy
		//allow us to execute the first attempt in 5 seconds and the last attend in 24 hours
		//which is useful if a user is executing the metrics without internet connection for several hours
		RetryAfter: func(attempt int) time.Duration {
			retryBacko := backo.NewBacko(retryBackoBaseDuration, retryBackoFactor, retryBackoJitter, retryBackoCap)
			return retryBacko.Duration(attempt)
		},
	}

	if shouldFlushQueueOnEachEvent {
		config.BatchSize = leastBatchSizeValue
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


func (segment *SegmentClient) TrackShouldSendMetricsUserElection(didUserAcceptSendingMetrics bool) error {

	newEvent, err := event.NewShouldSendMetricsUserElectionEvent(didUserAcceptSendingMetrics)
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

func (segment *SegmentClient) TrackLoadModule(moduleId, containerImage, serializedParams string) error {
	newEvent, err := event.NewLoadModuleEvent(moduleId, containerImage, serializedParams)
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

func (segment *SegmentClient) TrackExecuteModule(moduleId, serializedParams string) error {
	newEvent, err := event.NewExecuteModuleEvent(moduleId, serializedParams)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new execute module event")
	}

	if err := segment.track(newEvent); err != nil {
		return stacktrace.Propagate(err, "An error occurred tracking execute module event")
	}

	segment.client.Close()

	return nil
}

func (segment *SegmentClient) Close() (err error) {
	if err := segment.client.Close(); err != nil {
		return stacktrace.Propagate(err, "An error occurred closing the Segment client")
	}
	return nil
}

// ====================================================================================================
// 									   Private helper methods
// ====================================================================================================
func (segment *SegmentClient) track(event *event.Event) error {

	propertiesToTrack := analytics.NewProperties()

	eventProperties := event.GetProperties()

	for propertyKey, propertyValue := range eventProperties {
		propertiesToTrack.Set(propertyKey, propertyValue)
	}

	if err := segment.client.Enqueue(analytics.Track{
		Event:  event.GetName(),
		UserId: segment.userID,
		Context: segment.analyticsContext,
		Properties: propertiesToTrack,
	}); err != nil {
		return stacktrace.Propagate(err, "An error occurred enqueuing a new event in Segment client's queue")
	}
	return nil
}

func newAnalyticsContext(source metrics_source.Source, sourceVersion string) *analytics.Context {
	appInfo := analytics.AppInfo{
		Name: source.GetKey(),
		Version: sourceVersion,
	}

	analyticsContext := &analytics.Context{App: appInfo}

	return analyticsContext
}
