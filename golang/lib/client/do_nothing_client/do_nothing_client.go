package do_nothing_client

import "github.com/sirupsen/logrus"

//DoNothingClient: This metrics client implementation has been created for instantiate when user rejects
//sending metrics, so it doesn't really track metrics the only logic that it contains is loging
//the traking methods calls. It also can be used for test purpose
type DoNothingClient struct {

}

func NewDoNothingClient() *DoNothingClient {
	return &DoNothingClient{}
}

func (client *DoNothingClient) TrackShouldSendMetricsUserElection(didUserAcceptSendingMetrics bool) error {
	logrus.Debugf("Do-nothing metrics client TrackShouldSendMetricsUserElection called with argument didUserAcceptSendingMetrics '%v'; skipping sending event", didUserAcceptSendingMetrics)
	return nil
}

func (client *DoNothingClient) TrackCreateEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackCreateEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	return nil
}

func (client *DoNothingClient) TrackStopEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackStopEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	return nil
}

func (client *DoNothingClient) TrackDestroyEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackDestroyEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	return nil
}

func (client *DoNothingClient) TrackLoadModule(moduleId, containerImage, serializedParams string) error {
	logrus.Debugf("Do-nothing metrics client TrackLoadModule called with arguments moduleId '%v', containerImage '%v' and serializedParams '%v'; skipping sending event", moduleId, containerImage, serializedParams)
	return nil
}

func (client *DoNothingClient) TrackUnloadModule(moduleId string) error {
	logrus.Debugf("Do-nothing metrics client TrackUnloadModule called with argument moduleId '%v'; skipping sending event", moduleId)
	return nil
}

func (client *DoNothingClient) TrackExecuteModule(moduleId, serializedParams string) error {
	logrus.Debugf("Do-nothing metrics client TrackExecuteModule called with argument moduleId '%v' and serializedParams '%v'; skipping sending event", moduleId, serializedParams)
	return nil
}

func (client *DoNothingClient) Close() (err error) {
	logrus.Debugf("Do-nothing metrics client close method called")
	return nil
}
