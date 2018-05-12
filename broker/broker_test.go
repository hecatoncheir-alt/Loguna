package broker

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/hecatoncheir/loguna/configuration"
)

func TestBrokerCanSendMessageToNSQ(test *testing.T) {
	bro := New()

	config := configuration.New()
	err := bro.Connect(config.Development.Broker.Host, config.Development.Broker.Port)
	if err != nil {
		log.Println("Need started NSQ")
		log.Println(err)
	}

	item := map[string]string{"Name": "test item"}

	items, err := bro.ListenTopic(config.Development.LogunaTopic, config.APIVersion)
	if err != nil {
		test.Error(err)
	}

	err = bro.WriteToTopic(config.Development.LogunaTopic, item)
	if err != nil {
		test.Error(err)
	}

	defer bro.Producer.Stop()

	for item := range items {
		data := map[string]string{}
		json.Unmarshal(item, &data)
		if data["Name"] == "test item" {
			break
		}
	}
}
