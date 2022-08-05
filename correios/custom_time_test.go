package correios

import (
	"reflect"
	"testing"
	"time"
)

func TestCustomTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tr      CustomTime
		want    []byte
		wantErr bool
	}{
		{
			name:    "Valid 1",
			tr:      CustomTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			want:    []byte(`"2019-01-01T00:00:00"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonString string
		expected   time.Time
		wantErr    bool
	}{
		{
			name:       "Valid 1",
			jsonString: `"2019-01-01T00:00:00"`,
			expected:  time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		},
		{
			name:       "Valid 2",
			jsonString: `"2019-01-01T00:00:00"`,
			expected: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		}, {
			name:       "Invalid 1",
			jsonString: `"01/01/2019"`,
			wantErr:    true,
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &CustomTime{}
			if err := j.UnmarshalJSON([]byte(tt.jsonString)); (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}			
		})
	}
}
