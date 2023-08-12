package bluma

import (
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

type Attacker struct {
	path        Path
	methodIndex int
	frequency   int
	duration    time.Duration
}

func (atk *Attacker) Attack() {
	target := vegeta.Target{
		Method: atk.path.method[atk.methodIndex].method,
		URL:    atk.path.path,
	}
	targeter := vegeta.NewStaticTargeter(target)
	rate := vegeta.Rate{Freq: atk.frequency, Per: time.Second}

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	var results vegeta.Results
	reporter := vegeta.NewJSONReporter(&metrics)

	for res := range attacker.Attack(targeter, rate, atk.duration, "Vegeta Load Testing") {
		results.Add(res)
	}
	results.Close()
	reporter.Report(os.Stdout)

	fmt.Printf("Mean response time: %s\n", metrics.Latencies.Mean)
	fmt.Printf("99th percentile response time: %s\n", metrics.Latencies.P99)
}
