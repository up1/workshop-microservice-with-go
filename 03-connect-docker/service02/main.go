package main

import (
	"fmt"
	"service02/service"
)

var appName = "service02"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("9002")
}
