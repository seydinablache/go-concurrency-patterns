package main

import (
	"fmt"
	"github.com/configcat/go-sdk/v7"
)

func main() {
	client := configcat.NewClient("xxx") // <-- This is the actual SDK Key for your 'Production' environment.
	host := client.GetBoolValue("host", false, nil)
	isAwesomeFeatureEnabled := client.GetBoolValue("isAwesomeFeatureEnabled", false, nil)

	fmt.Println("host value from ConfigCat:", host)
	fmt.Println("isAwesomeFeatureEnabled value from ConfigCat:", isAwesomeFeatureEnabled)
}
