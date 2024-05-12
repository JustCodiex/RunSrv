package service

import (
	"encoding/json"
	"os"
)

type ServiceConfiguration struct {
}

func LoadConfiiguration(filepath string) (*ServiceConfiguration, error) {

	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	cfg := new(ServiceConfiguration)
	if err := json.Unmarshal(content, cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
