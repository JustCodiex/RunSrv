package service

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServiceConfiguration struct {
	ServicePort int // Port the hosted service will use for incoming/outgoing traffic
	HostPort    int // Port this service will use for incoming/outgoing traffic
	Deployment  *struct {
		Local  *LocalServiceConfig // Service configuration for a local service
		Docker *struct {
		}
	}
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

	if cfg.Deployment == nil || (cfg.Deployment.Docker == nil && cfg.Deployment.Local == nil) {
		return nil, fmt.Errorf("invalid configuration file - expected at least local deployment configuration")
	}

	if cfg.Deployment.Docker != nil && cfg.Deployment.Local != nil {
		return nil, fmt.Errorf("invalid configuration file - can only use local deployment or docker deployment")
	}

	if cfg.Deployment.Local != nil {
		if err := cfg.ValidateLocalDeployment(); err != nil {
			return nil, err
		}
	}

	return cfg, nil

}

func (cfg *ServiceConfiguration) ValidateLocalDeployment() error {

	local := cfg.Deployment.Local

	if len(local.Command) == 0 {
		return fmt.Errorf("invalid local deployment - expected non-empty startup command")
	}

	return nil

}

func (cfg *ServiceConfiguration) CreateHostedService() (HostedService, error) {
	if cfg.Deployment.Local != nil {
		hostedService := LocalService{
			cfg: cfg.Deployment.Local,
		}
		return hostedService, nil
	}
	return nil, fmt.Errorf("unknown error occured")
}
