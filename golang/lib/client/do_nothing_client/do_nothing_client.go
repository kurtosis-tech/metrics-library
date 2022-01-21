package do_nothing_client

import "github.com/sirupsen/logrus"

type DoNothingClient struct {

}

func NewDoNothingClient() *DoNothingClient {
	return &DoNothingClient{}
}

func (client DoNothingClient) TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error {
	logrus.Debugf("Do nothing client TrackUserAcceptSendingMetrics called")
	return nil
}

func (client DoNothingClient) TrackCreateEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackCreateEnclave called")
	return nil
}

func (client DoNothingClient) TrackStopEnclave() error {
	logrus.Debugf("Do nothing client TrackStopEnclave called")
	return nil
}

func (client DoNothingClient) TrackDestroyEnclave() error {
	logrus.Debugf("Do nothing client TrackDestroyEnclave called")
	return nil
}

func (client DoNothingClient) TrackCleanEnclave() error {
	logrus.Debugf("Do nothing client TrackCleanEnclave called")
	return nil
}
