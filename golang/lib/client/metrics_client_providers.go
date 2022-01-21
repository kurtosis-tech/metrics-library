package client

const (
	SnowPlow MetricsClientProvider = "snow-plow"
	Segment MetricsClientProvider = "segment"
	//It's used when users reject sending metrics
	DoNoting MetricsClientProvider = "do-noting"
)

type MetricsClientProvider string
