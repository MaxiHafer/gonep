package plant

import (
	"encoding/json"
	"github.com/maxihafer/gonep/pkg/gateway"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatus_UnmarshalJSON(t *testing.T) {
	var status Status
	err := json.Unmarshal(statusResponseData, &status)
	require.NoError(t, err)

	require.Equal(t, 0, status.CurrentWatts)
	require.Equal(t, 1748, status.TodayWattHours)
	require.Equal(t, 87317, status.YearWattHours)
	require.Equal(t, 89065, status.TotalWattHours)
	require.Equal(t, 96, status.KilogramsOfCO2Saved)
	require.Contains(t, status.Gateways, &gateway.Status{
		Id:                  "32c800c0",
		CurrentWatts:        0,
		TodayWattHours:      911,
		TotalWattHours:      45699,
		KilogramsOfCO2Saved: 49,
		Status:              "0000",
	})
	require.Contains(t, status.Gateways, &gateway.Status{
		Id:                  "32c833a0",
		CurrentWatts:        0,
		TodayWattHours:      837,
		TotalWattHours:      43366,
		KilogramsOfCO2Saved: 47,
		Status:              "0000",
	})
}
