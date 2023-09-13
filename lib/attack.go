package lib

import (
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func (atk *Attacker) Attack() {
	target := vegeta.Target{
		Method: atk.Path.method[atk.MethodIndex].method,
		URL:    atk.Path.path,
		Body:   atk.Body,
		Header: atk.Header,
	}
	targeter := vegeta.NewStaticTargeter(target)
	rate := vegeta.Rate{Freq: atk.Frequency, Per: time.Second}

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	var results vegeta.Results
	reporter := vegeta.NewJSONReporter(&metrics)

	for res := range attacker.Attack(targeter, rate, atk.Duration, "Vegeta Load Testing") {
		results.Add(res)
	}
	results.Close()
	reporter.Report(os.Stdout)

	fmt.Printf("Mean response time: %s\n", metrics.Latencies.Mean)
	fmt.Printf("99th percentile response time: %s\n", metrics.Latencies.P99)
}
