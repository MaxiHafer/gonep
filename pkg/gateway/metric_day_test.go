package gateway

import (
	"encoding/json"
	"github.com/maxihafer/gonep/pkg/pointer"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDayMetric_UnmarshalJSON(t *testing.T) {
	var metrics []*dayMetric

	err := json.Unmarshal(dayResponseBody, &metrics)
	require.NoError(t, err)

	require.Contains(t, metrics, &dayMetric{
		ts: pointer.Of(
			time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		),
		kwh: 0.3,
	})

	require.Contains(t, metrics, &dayMetric{
		ts: pointer.Of(time.Date(
			2023, time.January, 2, 0, 0, 0, 0, time.Local),
		),
		kwh: 1.8,
	})
}
