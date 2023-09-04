package gateway

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMonthMetric_UnmarshalJSON(t *testing.T) {
	var metrics []*monthMetric

	err := json.Unmarshal(monthResponseBody, &metrics)
	require.NoError(t, err)

	require.Contains(t, metrics, &monthMetric{
		Timestamp: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		Watts:     29,
	})

	require.Contains(t, metrics, &monthMetric{
		Timestamp: time.Date(2023, time.February, 1, 0, 0, 0, 0, time.Local),
		Watts:     23,
	})
}
