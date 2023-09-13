package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttack(t *testing.T) {
	t.Run("Attack", func(t *testing.T) {
		t.Skip()
		atk := Attacker{
			path: Path{
				path: "http://localhost:8080/health",
				method: []Method{
					{
						method: "GET",
					},
				},
			},
			methodIndex: 0,
			frequency:   10,
			duration:    1,
		}
		assert.NotPanics(t, atk.Attack)
	})
}
