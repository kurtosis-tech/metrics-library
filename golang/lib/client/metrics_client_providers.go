package client

const (
	SnowPlow MetricsClientProvider = "snow-plow"
	Segment MetricsClientProvider = "segment"
)

type MetricsClientProvider string
