package do_nothing_client

import "github.com/sirupsen/logrus"

type DoNothingClient struct {

}

func NewDoNothingClient() *DoNothingClient {
	return &DoNothingClient{}
}

func (client *DoNothingClient) TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error {
	logrus.Debugf("Do nothing client TrackUserAcceptSendingMetrics called")
	return nil
}

func (client *DoNothingClient) TrackCreateEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackCreateEnclave called")
	return nil
}

func (client *DoNothingClient) TrackStopEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackStopEnclave called")
	return nil
}

func (client *DoNothingClient) TrackDestroyEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackDestroyEnclave called")
	return nil
}

func (client *DoNothingClient) TrackCleanEnclave(shouldCleanAll bool) error {
	logrus.Debugf("Do nothing client TrackCleanEnclave called")
	return nil
}

func (client *DoNothingClient) TrackLoadModule(moduleId string) error {
	logrus.Debugf("Do nothing client TrackLoadModule called")
	return nil
}

func (client *DoNothingClient) TrackUnloadModule(moduleId string) error {
	logrus.Debugf("Do nothing client TrackUnloadModule called")
	return nil
}

func (client *DoNothingClient) TrackExecuteModule(moduleId string) error {
	logrus.Debugf("Do nothing client TrackExecuteModule called")
	return nil
}
