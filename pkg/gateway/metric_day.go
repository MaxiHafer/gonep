package gateway

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	dayFormat = "01/02"
)

var _ Metric = (*dayMetric)(nil)
var _ json.Unmarshaler = (*dayMetric)(nil)

type dayMetric struct {
	Timestamp *time.Time
	Watts     int
}

func (d *dayMetric) UnmarshalJSON(bytes []byte) error {
	var data []interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if len(data) != 2 {
		return fmt.Errorf("error while unmarshaling metric: malformed datapoint: %v", data)
	}

	timestampStr := fmt.Sprintf("%v", data[0])

	wattsF, ok := data[1].(float64)
	if !ok {
		return fmt.Errorf("%v is not representable as type float64", data[1])
	}

	ts, err := time.Parse(dayFormat, timestampStr)
	if err != nil {
		return fmt.Errorf("timestamp string: %s does not match time format: %s", timestampStr, dayFormat)
	}

	now := time.Now()
	ts = time.Date(now.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.Local)

	d.Timestamp = &ts
	d.Watts = int(wattsF)

	return nil
}

func (d *dayMetric) Time() *time.Time {
	return d.Timestamp
}

func (d *dayMetric) Value() int {
	return d.Watts
}
