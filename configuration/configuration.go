package configuration

import (
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	APIVersion string

	Production struct {
		LogunaTopic string
		LogFilePath string

		Broker struct {
			Host string
			Port int
		}
	}

	Development struct {
		LogunaTopic string
		LogFilePath string

		Broker struct {
			Host string
			Port int
		}
	}
}

func New() *Configuration {
	configuration := Configuration{}

	apiVersion := os.Getenv("Api-Version")
	if apiVersion == "" {
		configuration.APIVersion = "1.0.0"
	} else {
		configuration.APIVersion = apiVersion
	}

	productionBrokerHostFromEnvironment := os.Getenv("Production-Broker-Host")
	if productionBrokerHostFromEnvironment == "" {
		configuration.Production.Broker.Host = "192.168.99.100"
	} else {
		configuration.Production.Broker.Host = productionBrokerHostFromEnvironment
	}

	productionBrokerPortFromEnvironment := os.Getenv("Production-Broker-Port")
	if productionBrokerPortFromEnvironment == "" {
		configuration.Production.Broker.Port = 4150
	} else {
		port, err := strconv.Atoi(productionBrokerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Production.Broker.Port = port
	}

	developmentBrokerHostFromEnvironment := os.Getenv("Development-Broker-Host")
	if developmentBrokerHostFromEnvironment == "" {
		configuration.Development.Broker.Host = "192.168.99.100"
	} else {
		configuration.Development.Broker.Host = developmentBrokerHostFromEnvironment
	}

	developmentBrokerPortFromEnvironment := os.Getenv("Development-Broker-Port")
	if developmentBrokerPortFromEnvironment == "" {
		configuration.Development.Broker.Port = 4150
	} else {
		port, err := strconv.Atoi(developmentBrokerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Development.Broker.Port = port
	}

	productionLogunaTopic := os.Getenv("Production-Loguna-Topic")
	if productionLogunaTopic == "" {
		configuration.Production.LogunaTopic = "Loguna"
	} else {
		configuration.Production.LogunaTopic = productionLogunaTopic
	}

	developmentLogunaTopic := os.Getenv("Development-Loguna-Topic")
	if developmentLogunaTopic == "" {
		configuration.Development.LogunaTopic = "DevLoguna"
	} else {
		configuration.Development.LogunaTopic = developmentLogunaTopic
	}

	productionLogFilePath := os.Getenv("Production-Log-File-Path")
	if productionLogFilePath == "" {
		configuration.Production.LogFilePath = "log"
	} else {
		configuration.Production.LogFilePath = productionLogFilePath
	}

	developmentLogFilePath := os.Getenv("Development-Log-File-Path")
	if developmentLogFilePath == "" {
		configuration.Development.LogFilePath = "dev_log"
	} else {
		configuration.Development.LogFilePath = developmentLogFilePath
	}

	return &configuration
}
