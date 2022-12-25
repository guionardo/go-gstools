package schedule

import (
	"testing"
	"time"
)

func TestNewScheduleEmptyShouldHaveNoNext(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		s := NewSchedule("test")
		if s.NextRun().IsZero() == false {
			t.Errorf("NewSchedule() = %v, want %v", s.NextRun(), time.Time{})
		}
	})
}

func TestNewSchedule(t *testing.T) {

	tests := []struct {
		name     string
		schedule *Schedule
		valid    func(s *Schedule) bool
	}{
		{
			name:     "Default",
			schedule: NewSchedule("default"),
			valid: func(s *Schedule) bool {
				return s.Name == "default"
			},
		},
		{
			name:     "Every",
			schedule: NewSchedule("every").Every(1 * time.Second),
			valid: func(s *Schedule) bool {
				return s.runsEvery == 1*time.Second
			},
		},
		{
			name:     "After",
			schedule: NewSchedule("after").After(time.Now().Add(1 * time.Hour)),
			valid: func(s *Schedule) bool {
				return s.runsAfter.After(time.Now().Add(59 * time.Minute))
			},
		},
		{
			name:     "Disabled",
			schedule: NewSchedule("disabled").Enabled(false),
			valid: func(s *Schedule) bool {
				return s.NextRun().IsZero()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.valid(tt.schedule) {
				t.Errorf("NewSchedule() = %v", tt.schedule)
			}
		})
	}
}
