package correios

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

type CustomTime time.Time

var Formats = mapset.NewSet[string]()
var DefaultFormat = "2006-01-02T15:04:05"

func init() {
	Formats.Add(DefaultFormat)
	Formats.Add("2006-01-02 15:04:05")
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	for format := range Formats.Iterator().C {
		body, err := json.Marshal(time.Time(t).Format(format))
		if err == nil {
			return body, nil
		}
	}
	return nil, fmt.Errorf("unsupported CustomTime format")
}

func (j *CustomTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	t, err := time.Parse(DefaultFormat, s)
	if err != nil {
		return err
	}
	*j = CustomTime(t)
	return nil
}
