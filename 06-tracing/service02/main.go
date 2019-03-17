package main

import (
	"fmt"
	"service02/service"
	"service02/tracing"
)

var appName = "service02"

func main() {
	initializeTracing()
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("9002")
}

func initializeTracing() {
	tracing.InitTracing("http://localhost:9411", appName)
}
