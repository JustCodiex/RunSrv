package main

import (
	"flag"
	"fmt"

	service "runsrv.com/runsrv/Service"
)

func main() {

	serviceConfigFilepath := flag.String("config", "app.config.json", "Specifies the configuration file to load")
	flag.Parse()

	_, err := service.LoadConfiiguration(*serviceConfigFilepath)
	if err != nil {
		fmt.Printf("Failed loading service configuration:\n\t%s\n", err.Error())
		return
	}

}
