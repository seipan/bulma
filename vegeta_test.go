package bluma

import "testing"

func TestAttack(t *testing.T) {
	t.Run("Attack", func(t *testing.T) {
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
		atk.Attack()
	})
}
