package main

import (
	"fmt"
	"github.com/hecatoncheir/loguna/broker"
	"github.com/hecatoncheir/loguna/configuration"
)

func main() {
	config := configuration.New()

	bro := broker.New()
	topicEvents, err := bro.ListenTopic(
		config.Production.LogunaTopic, config.APIVersion)
	if err != nil {
		panic(err)
	}

	for event := range topicEvents {
		fmt.Println(string(event))
	}
}
