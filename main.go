package main

import (
	"flag"
	"fmt"

	service "runsrv.com/runsrv/Service"
)

func main() {

	serviceConfigFilepath := flag.String("config", "app.config.json", "Specifies the configuration file to load")
	flag.Parse()

	fmt.Printf("Loading configuration file from %s\n", *serviceConfigFilepath)
	cfg, err := service.LoadConfiiguration(*serviceConfigFilepath)
	if err != nil {
		fmt.Printf("Failed loading service configuration:\n\t%s\n", err.Error())
		return
	}

	fmt.Println("Successfully loaded configuration file")
	fmt.Printf("Loaded service port: %v\n", cfg.ServicePort)

	hostedService, err := cfg.CreateHostedService()
	if err != nil {
		fmt.Printf("Failed creating service:\n\t%s\n", err.Error())
		return
	}

	if err := hostedService.Start(); err != nil {
		fmt.Printf("Failed starting service:\n\t%s\n", err.Error())
		return
	}

}
