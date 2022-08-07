package shared

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestSaveOutput(t *testing.T) {
	type (
		data struct {
			Name string
		}
		args struct {
			data           data
			output_type    string
			output_destiny string
			content        []byte
		}
	)
	testData := data{Name: "John"}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "JSON console",
			args: args{
				data:           testData,
				output_type:    "json",
				output_destiny: "console",
			},
			wantErr: false,
		}, {
			name: "YAML console",
			args: args{
				data:           testData,
				output_type:    "yaml",
				output_destiny: "console",
			},
		},
		{
			name: "TOML console",
			args: args{
				data:           testData,
				output_type:    "toml",
				output_destiny: "console",
			},
		},
		{
			name: "Invalid output type",
			args: args{
				data:           testData,
				output_type:    "invalid",
				output_destiny: "console",
			},
			wantErr: true,
		},
		{
			name: "JSON file",
			args: args{
				data:           testData,
				output_type:    "json",
				output_destiny: path.Join(t.TempDir(), "data.json"),
				content: []byte(`{
  "Name": "John"
}`),
			},
			wantErr: false,
		}, {
			name: "YAML file",
			args: args{
				data:           testData,
				output_type:    "yaml",
				output_destiny: path.Join(t.TempDir(), "data.yaml"),
				content:        []byte("name: John\n"),
			},
		},
		{
			name: "TOML file",
			args: args{
				data:           testData,
				output_type:    "toml",
				output_destiny: path.Join(t.TempDir(), "data.toml"),
				content:        []byte("Name = \"John\"\n"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveOutput(tt.args.data, tt.args.output_type, tt.args.output_destiny); (err != nil) != tt.wantErr {
				t.Errorf("SaveOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.output_destiny != "console" {
				content, err := os.ReadFile(tt.args.output_destiny)
				if err != nil {
					t.Errorf("SaveOutput() error reading file %s = %v", tt.args.output_destiny, err)
				}
				if !reflect.DeepEqual(content, tt.args.content) {
					t.Errorf("SaveOutput() content = %s, want %s", content, tt.args.content)
				}
			}
		})
	}
}
