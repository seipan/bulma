package bluma

import (
	"context"
	"testing"
)

func TestParseOpenAPI(t *testing.T) {
	t.Run("ParseOpenAPI", func(t *testing.T) {
		t.Skip()
		o := NewOpenAPI("testdata/openapi.yaml")
		o.Parse(context.Background())
	})
}
