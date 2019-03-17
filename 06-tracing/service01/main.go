package main

import (
	"fmt"
	"service01/service"
	"service01/tracing"
)

var appName = "service01"

func main() {
	initializeTracing()
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("9001")
}

func initializeTracing() {
	tracing.InitTracing("http://localhost:9411", appName)
}
