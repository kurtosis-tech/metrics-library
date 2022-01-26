package do_nothing_client

import "github.com/sirupsen/logrus"

type DoNothingClient struct {

}

func NewDoNothingClient() *DoNothingClient {
	return &DoNothingClient{}
}

func (client *DoNothingClient) TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error {
	logrus.Debugf("Do nothing client TrackUserAcceptSendingMetrics called with argument userAcceptSendingMetrics '%v'", userAcceptSendingMetrics)
	return nil
}

func (client *DoNothingClient) TrackCreateEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackCreateEnclave called with argument enclaveId '%v'", enclaveId)
	return nil
}

func (client *DoNothingClient) TrackStopEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackStopEnclave called with argument enclaveId '%v'", enclaveId)
	return nil
}

func (client *DoNothingClient) TrackDestroyEnclave(enclaveId string) error {
	logrus.Debugf("Do nothing client TrackDestroyEnclave called with argument enclaveId '%v'", enclaveId)
	return nil
}

func (client *DoNothingClient) TrackCleanEnclave(shouldCleanAll bool) error {
	logrus.Debugf("Do nothing client TrackCleanEnclave called with argument shouldCleanAll '%v'", shouldCleanAll)
	return nil
}

func (client *DoNothingClient) TrackLoadModule(moduleId string) error {
	logrus.Debugf("Do nothing client TrackLoadModule called with argument moduleId '%v'", moduleId)
	return nil
}

func (client *DoNothingClient) TrackUnloadModule(moduleId string) error {
	logrus.Debugf("Do nothing client TrackUnloadModule called with argument moduleId '%v'", moduleId)
	return nil
}

func (client *DoNothingClient) TrackExecuteModule(moduleId string) error {
	logrus.Debugf("Do nothing client TrackExecuteModule called with argument moduleId '%v'", moduleId)
	return nil
}
