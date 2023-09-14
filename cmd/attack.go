package cmd

import (
	"context"
	"fmt"

	"github.com/seipan/bulma/lib"
)

func ParseAndAttack(ctx context.Context, path string) error {
	oapi := lib.NewOpenAPI(path)
	paths, err := oapi.Parse(ctx)
	if err != nil {
		return fmt.Errorf("failed to parse openapi: %w", err)
	}
	atks, err := ParthOpenAPItoAttacker(paths)
	if err != nil {
		return fmt.Errorf("failed to convert openapi to attacker: %w", err)
	}
	for _, atk := range atks {
		atk.Attack()
	}
	return nil
}

func ParthOpenAPItoAttacker(pathes []lib.Path) ([]lib.Attacker, error) {
	var res []lib.Attacker
	for i, path := range pathes {
		mtd := path.Method(i)
		body := mtd.Bodys()
		bodyjson, err := body[0].MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal json: %w", err)
		}
		atk := lib.Attacker{
			Path:        path,
			MethodIndex: i,
			Body:        bodyjson,
		}
		res = append(res, atk)
	}
	return res, nil
}
