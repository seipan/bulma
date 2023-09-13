package lib

import (
	"net/http"
	"time"
)

type Attacker struct {
	Path        Path
	MethodIndex int
	Body        []byte
	Header      http.Header
	Frequency   int
	Duration    time.Duration
}
