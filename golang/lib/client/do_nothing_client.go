package client

import "github.com/sirupsen/logrus"

// doNothingClient: This metrics client implementation has been created for instantiate when user rejects
// sending metrics, so it doesn't really track metrics the only logic that it contains is loging
// the traking methods calls. It also can be used for test purpose
type doNothingClient struct {
	callback Callback
}

func newDoNothingClient(callback Callback) *doNothingClient {
	return &doNothingClient{callback: callback}
}

func (client *doNothingClient) TrackShouldSendMetricsUserElection(didUserAcceptSendingMetrics bool) error {
	logrus.Debugf("Do-nothing metrics client TrackShouldSendMetricsUserElection called with argument didUserAcceptSendingMetrics '%v'; skipping sending event", didUserAcceptSendingMetrics)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackCreateEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackCreateEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackStopEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackStopEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackDestroyEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackDestroyEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackLoadModule(moduleId, containerImage, serializedParams string) error {
	logrus.Debugf("Do-nothing metrics client TrackLoadModule called with arguments moduleId '%v', containerImage '%v' and serializedParams '%v'; skipping sending event", moduleId, containerImage, serializedParams)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackUnloadModule(moduleId string) error {
	logrus.Debugf("Do-nothing metrics client TrackUnloadModule called with argument moduleId '%v'; skipping sending event", moduleId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackExecuteModule(moduleId, serializedParams string) error {
	logrus.Debugf("Do-nothing metrics client TrackExecuteModule called with argument moduleId '%v' and serializedParams '%v'; skipping sending event", moduleId, serializedParams)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackRunStarlarkPackage(packageId string, serializedArgs string, isRemote bool, isDryRun bool) error {
	logrus.Debugf("Do-nothing metrics client TrackRunStarlarkPackage called with arguments packageId '%v', serializedArgs '%v', isRemote '%v' and isDryrun '%v'; skipping sending event", isRemote, packageId, serializedArgs, isDryRun)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackRunStarlarkScript(serializedScript string, serializedArgs string, isDryRun bool) error {
	logrus.Debugf("Do-nothing metrics client TrackRunStarlarkScript with arguments serializedScript '%v', serializedArgs '%v' and isDryRun '%v'; skipping sending event", serializedScript, serializedArgs, isDryRun)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) close() (err error) {
	logrus.Debugf("Do-nothing metrics client close method called")
	return nil
}
