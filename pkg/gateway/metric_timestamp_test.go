package gateway

import (
	"encoding/json"
	"github.com/maxihafer/gonep/pkg/pointer"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMetric_UnmarshalJSON(t *testing.T) {
	var metrics []*timestampMetric
	err := json.Unmarshal(timestampResponseBody, &metrics)

	require.NoError(t, err)
	require.Contains(t, metrics, &timestampMetric{
		ts: pointer.Of(
			time.UnixMilli(1693161960000),
		),
		watts: 29,
	})
	require.Contains(t, metrics, &timestampMetric{
		ts: pointer.Of(
			time.UnixMilli(1693162260000),
		),
		watts: 23,
	})
}
