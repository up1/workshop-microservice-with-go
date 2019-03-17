package main

import (
	"fmt"
	"log"
	"os"
	"service02/service"

	consulapi "github.com/hashicorp/consul/api"
)

var appName = "service02"

func main() {
	registerServiceWithConsul()
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("9002")
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "service02"
	registration.Name = "service02"
	registration.Tags = []string{"urlprefix-/service02"}
	address := hostname()
	registration.Address = address
	registration.Port = 9002
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:9002/healthcheck", address)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
