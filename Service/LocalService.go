package service

import (
	"fmt"
	"os/exec"
)

type LocalServiceConfig struct {
	Command       string   // The startup command
	Arguments     []string // The arguments to pass onto the startup command
	WorkDirectory string   // The working directory of the application
}

type LocalService struct {
	cfg *LocalServiceConfig
}

func (s LocalService) Start() error {

	cmd := exec.Command(s.cfg.Command, s.cfg.Arguments...)
	if len(s.cfg.WorkDirectory) > 0 {
		cmd.Dir = s.cfg.WorkDirectory
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	if cmd.ProcessState.Exited() {
		return fmt.Errorf("failed starting local service - exit code = %v", cmd.ProcessState.ExitCode())
	}

	return nil
}
