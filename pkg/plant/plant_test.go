package plant

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlantList_UnmarshalJSON(t *testing.T) {
	var plants PlantList
	err := json.Unmarshal(listResponseData, &plants)
	require.NoError(t, err)

	require.Equal(t, &Plant{
		Sid:               "TEST_20230827_xY47",
		UserEmail:         "testuser@example.com",
		SiteName:          "Test Home",
		InstallationEmail: "installer@example.com",
		Country:           "US",
		State:             "California",
		City:              "Los Angeles",
		Street:            "123 Main St",
		Latitude:          "34.052235",
		Longitude:         "-118.243683",
		ImageRef:          "http://example.com/testlogo.jpg",
		GroupStrSN:        "aBcDeF12,GhIjKl34",
		TemperatureUnit:   "Celsius",
		Gateways: []string{
			"aBcDeF12",
			"GhIjKl34",
		},
	}, plants[0])
}
