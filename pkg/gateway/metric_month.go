package gateway

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	monthFormat = "2006.01"
)

var _ Metric = (*monthMetric)(nil)
var _ json.Unmarshaler = (*monthMetric)(nil)

type monthMetric struct {
	ts  *time.Time
	kwh float64
}

func (d *monthMetric) UnmarshalJSON(bytes []byte) error {
	var data []interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if len(data) != 2 {
		return fmt.Errorf("error while unmarshaling metric: malformed datapoint: %v", data)
	}

	timestampStr := fmt.Sprintf("%v", data[0])

	kwh, ok := data[1].(float64)
	if !ok {
		return fmt.Errorf("%v is not convertible to type string float64", data[1])
	}

	ts, err := time.ParseInLocation(monthFormat, timestampStr, time.Local)
	if err != nil {
		return fmt.Errorf("timestamp string: %s does not match time format: %s", timestampStr, monthFormat)
	}

	d.ts = &ts
	d.kwh = kwh

	return nil
}

func (d *monthMetric) Time() *time.Time {
	return d.ts
}

func (d *monthMetric) KilowattHours() float64 {
	return d.kwh
}
