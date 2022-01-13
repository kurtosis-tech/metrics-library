package main

import (
	sp "github.com/snowplow/snowplow-golang-tracker/v2/tracker"
	"log"
	"time"
)

func main() {









}


func runSnowPlotTrialTracker() {
	subject := sp.InitSubject()
	emitter := sp.InitEmitter(
		sp.RequireCollectorUri("8e280f93-12b7-4610-bd19-a5d7bc9e41dd.app.try-snowplow.com"),
		sp.OptionCallback(func(g []sp.CallbackResult, b []sp.CallbackResult) {
			log.Println("Successes: " + sp.IntToString(len(g)))
			log.Println("Failures: " + sp.IntToString(len(b)))
			for _, val := range b {
				log.Println("Count: " + sp.IntToString(val.Count))
				log.Println("Status: " + sp.IntToString(val.Status))
			}
		}),
		sp.OptionProtocol("https"),
		)
	tracker := sp.InitTracker(
		sp.RequireEmitter(emitter),
		sp.OptionSubject(subject),
		sp.OptionNamespace("kurtosistech"),
		sp.OptionAppId("kurtosis-cli"),
		sp.OptionPlatform("pc"), //Available values https://github.com/snowplow/enrich/issues/450
	)

	tracker.TrackStructEvent(sp.StructuredEvent{
		Category: sp.NewString("CLI"),
		Action: sp.NewString("create"),
		Property: sp.NewString("enclave"),
		Label: sp.NewString("kurtosis-engine"),
	})

	tracker.TrackStructEvent(sp.StructuredEvent{
		Category: sp.NewString("CLI"),
		Action: sp.NewString("stop"),
		Property: sp.NewString("enclave"),
		Label: sp.NewString("kurtosis-engine"),
	})

	tracker.TrackStructEvent(sp.StructuredEvent{
		Category: sp.NewString("CLI"),
		Action: sp.NewString("remove"),
		Property: sp.NewString("enclave"),
		Label: sp.NewString("kurtosis-engine"),
	})

	tracker.TrackStructEvent(sp.StructuredEvent{
		Category: sp.NewString("CLI"),
		Action: sp.NewString("load"),
		Property: sp.NewString("module"),
		Label: sp.NewString("kurtosis-tech/eth2-merge-kurtosis-module"),
	})

	tracker.TrackStructEvent(sp.StructuredEvent{
		Category: sp.NewString("CLI"),
		Action: sp.NewString("unload"),
		Property: sp.NewString("module"),
		Label: sp.NewString("kurtosis-tech/eth2-merge-kurtosis-module"),
	})


	time.Sleep(5 * time.Minute)
}