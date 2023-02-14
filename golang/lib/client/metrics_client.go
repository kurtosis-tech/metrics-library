package client

type MetricsClient interface {
	TrackShouldSendMetricsUserElection(didUserAcceptSendingMetrics bool) error
	TrackCreateEnclave(enclaveId string) error
	TrackStopEnclave(enclaveId string) error
	TrackDestroyEnclave(enclaveId string) error
	TrackKurtosisRun(packageId string, isRemote bool, isDryRun bool, isScript bool) error
	close() (err error)
}
