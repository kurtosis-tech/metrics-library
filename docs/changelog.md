# TBD

# 0.1.1
### Features
* Added `MetricsClient` interface to define Kurtosis metrics abstraction behaviour
* Added `SegmentClient` implementation of the `MetricsClient` using Segment provider
* Added `DoNothingClient` implementation of the `MetricsClient` used when users decide reject to send metrics
* Added `Event` object to set the fields involve in a Kurtosis Event. `Category` and `Action` fields are mandatory
* Added event types to centralize Kurtosis events data
* Added `Source` type to define the metrics application source
* Added metrics client creator func to create default metrics type depending on the passed arguments

# 0.1.0
* Initial commit
