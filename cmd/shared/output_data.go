package shared

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

func SaveOutput(data any, output_type string, output_destiny string) error {
	var (
		content []byte
		err     error
	)
	switch output_type {
	case "json":
		content, err = json.MarshalIndent(data, "", "  ")
	case "yaml":
		content, err = yaml.Marshal(data)
	case "toml":
		content, err = toml.Marshal(data)
	default:
		err = fmt.Errorf("Formato de saída inválido [%s]", output_type)
	}
	if err != nil {
		return err
	}
	if output_destiny == "console" {
		fmt.Print(string(content))
		return nil
	}
	file, err := os.Create(output_destiny)
	if err == nil {
		defer file.Close()
		_, err = file.Write(content)
	}
	return err
}
