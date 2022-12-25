package schedule

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		got := Today()
		if got.Hour() != 0 || got.Minute() != 0 || got.Second() != 0 || got.Nanosecond() != 0 {
			t.Errorf("Today() = %v, want %v", got, time.Time{})
		}
	})

}

func TestWallClock(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		if got := WallClock(); got.Seconds() == 0 {
			t.Errorf("WallClock() = %v, want positive", got)
		}
	})

}

func TestTomorrow(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		if got := Tomorrow(); got.Day() == time.Now().Day() {
			t.Errorf("Tomorrow() = %v", got)
		}
	})

}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		wantDuration time.Duration
		wantErr      bool
	}{
		{
			name:         "Duration",
			str:          "1h30m45s",
			wantDuration: 90*time.Minute + 45*time.Second,
			wantErr:      false,
		},
		{
			name:         "Hour/Minute",
			str:          "01:30",
			wantDuration: 90 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "Hour/Minute/Second",
			str:          "01:30:45",
			wantDuration: 90*time.Minute + 45*time.Second,
			wantErr:      false,
		},
		{
			name:    "Invalid",
			str:     "invalid",
			wantErr: true,
		},
		{
			name:    "Invalid hour",
			str:     "x:00:00",
			wantErr: true,
		},
		{
			name:    "Invalid hour 2",
			str:     "25:00:00",
			wantErr: true,
		},
		{
			name:    "Invalid minute",
			str:     "00:x:00",
			wantErr: true,
		},
		{
			name:    "Invalid second",
			str:     "00:00:61",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDuration, err := ParseDuration(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDuration != tt.wantDuration {
				t.Errorf("ParseDuration() = %v, want %v", gotDuration, tt.wantDuration)
			}
		})
	}
}
