package gateway

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMetric_UnmarshalJSON(t *testing.T) {
	var metrics []*dayMetric
	err := json.Unmarshal(timestampResponseBody, &metrics)

	require.NoError(t, err)
	require.Contains(t, metrics, &dayMetric{
		Timestamp: time.UnixMilli(1693161960000),
		Watts:     29,
	})
	require.Contains(t, metrics, &dayMetric{
		Timestamp: time.UnixMilli(1693162260000),
		Watts:     23,
	})
}
