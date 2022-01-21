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
