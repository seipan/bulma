package lib

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOpenAPI(t *testing.T) {
	t.Run("ParseOpenAPI", func(t *testing.T) {
		o := NewOpenAPI("testdata/openapi.json")
		res, err := o.Parse(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, "GET", res[0].method[0].method)
		assert.Equal(t, "/health", res[0].Path())
	})
}
