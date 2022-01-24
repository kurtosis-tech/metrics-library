# TBD
### Features
* Added `MetricsClient` interface to define Kurtosis metrics abstraction behaviour
* Added `SnowPlowClient` implementation of the `MetricsClient` using SnowPlow provider
* Added `Event` object to set the fields involve in a Kurtosis Event. `Category` and `Action` fields are mandatory
* Added event types to centralize Kurtosis events data
* Added `Source` type to define the metrics application source
* Added metrics client creator func to create default metrics provider or and specific metrics provider
* Added `DoNothingClient` which is and implementation of the `MetricsClient` used when users decide reject send metrics

# 0.1.0
* Initial commit
