package version

import (
	"reflect"
	"testing"
)

func TestVersionParse(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    Version
		wantErr bool
	}{
		{
			name:    "Valid version 1",
			version: "1.2.3",
			want:    Version{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "Valid version 2",
			version: "v4.5.6",
			want:    Version{4, 5, 6},
			wantErr: false,
		},
		{
			name:    "Invalid version",
			version: "a.3",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VersionParse(tt.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionParse(%s) error = %v, wantErr %v", tt.version, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VersionParse(%s) = %v, want %v", tt.version, got, tt.want)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{
			name:     "Compare 1.2.3 and 1.2.3",
			version1: "1.2.3",
			version2: "1.2.3",
			want:     0,
		},
		{
			name:     "Compare 0.0.1 and 0.0.2",
			version1: "0.0.1",
			version2: "0.0.2",
			want:     -1,
		},
		{
			name:     "Compare 1.2.3 and 0.2.2",
			version1: "1.2.3",
			version2: "0.2.2",
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, _ := VersionParse(tt.version1)
			v2, _ := VersionParse(tt.version2)

			if got := v1.Compare(v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
