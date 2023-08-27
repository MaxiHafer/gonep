package gateway

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	metricData = []byte(`[
	  [
		1693161960000,
		29
	  ],
	  [
		1693162260000,
		23
	  ]
	]`)
)

func TestMetric_UnmarshalJSON(t *testing.T) {
	var metrics []*Metric
	err := json.Unmarshal(metricData, &metrics)

	require.NoError(t, err)
	require.Contains(t, metrics, &Metric{
		Timestamp: time.UnixMilli(1693161960000),
		Watts:     29,
	})
	require.Contains(t, metrics, &Metric{
		Timestamp: time.UnixMilli(1693162260000),
		Watts:     23,
	})
}
