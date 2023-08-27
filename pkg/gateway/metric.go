package gateway

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ json.Unmarshaler = (*Metric)(nil)

type Metric struct {
	Timestamp time.Time
	Watts     int
}

func (d *Metric) UnmarshalJSON(bytes []byte) error {
	var data []int64
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if len(data) != 2 {
		return fmt.Errorf("error while unmarshaling metric: malformed datapoint: %v", data)
	}

	d.Timestamp = time.UnixMilli(data[0])
	d.Watts = int(data[1])

	return nil
}
