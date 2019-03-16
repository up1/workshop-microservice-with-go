package main

import (
	"fmt"
	"service01/service"
)

var appName = "service01"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("9001")
}
