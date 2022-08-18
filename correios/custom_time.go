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
	return json.Marshal(time.Time(t).Format(DefaultFormat))
}

func (j *CustomTime) UnmarshalJSON(data []byte) error {
	for format := range Formats.Iterator().C {
		s := strings.Trim(string(data), "\"")
		t, err := time.Parse(format, s)
		if err == nil {
			*j = CustomTime(t)
			return nil
		}
	}
	return fmt.Errorf("unsupported CustomTime format")
}

func (j *CustomTime) ToTime() time.Time {
	return time.Time(*j)
}
