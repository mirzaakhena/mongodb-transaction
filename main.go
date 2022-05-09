package main

import (
	"flag"
	"fmt"
	"mongodb-trx/application"
	"mongodb-trx/shared/driver"
)

func main() {
	appMap := map[string]func() driver.RegistryContract{
		"mirza": application.NewMirza(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
