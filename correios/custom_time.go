package correios

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomTime time.Time

const format = "2006-01-02T15:04:05"

func (t CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(format))
}

func (j *CustomTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	*j = CustomTime(t)
	return nil
}
