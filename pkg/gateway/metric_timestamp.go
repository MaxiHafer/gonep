package gateway

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ Metric = (*timestampMetric)(nil)
var _ json.Unmarshaler = (*timestampMetric)(nil)

type timestampMetric struct {
	Timestamp time.Time
	Watts     int
}

func (d *timestampMetric) UnmarshalJSON(bytes []byte) error {
	var data []interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if len(data) != 2 {
		return fmt.Errorf("error while unmarshaling metric: malformed datapoint: %v", data)
	}

	timestampF, ok := data[0].(float64)
	if !ok {
		return fmt.Errorf("%v is not representable as type float64", data[0])
	}

	wattsF, ok := data[1].(float64)
	if !ok {
		return fmt.Errorf("%v is not representable as type float64", data[1])
	}

	d.Timestamp = time.UnixMilli(int64(timestampF))
	d.Watts = int(wattsF)

	return nil
}

func (d *timestampMetric) Time() time.Time {
	return d.Timestamp
}

func (d *timestampMetric) Value() int {
	return d.Watts
}
