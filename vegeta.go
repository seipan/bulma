package bluma

import (
	"net/http"
	"time"
)

type Attacker struct {
	path        Path
	methodIndex int
	body        []byte
	header      http.Header
	frequency   int
	duration    time.Duration
}
