package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttack(t *testing.T) {
	t.Run("Attack", func(t *testing.T) {
		t.Skip()
		atk := Attacker{
			Path: Path{
				path: "http://localhost:8080/health",
				method: []Method{
					{
						method: "GET",
					},
				},
			},
			MethodIndex: 0,
			Frequency:   10,
			Duration:    1,
		}
		assert.NotPanics(t, atk.Attack)
	})
}
