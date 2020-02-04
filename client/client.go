package main

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	mc mqtt.Client
	o  Opts
}

// NewClient returns a new paho mqtt client configured by passed options
func NewClient(o Opts) Client {
	opts := mqtt.NewClientOptions().AddBroker(
		o.Broker).SetClientID(o.ID)

	log.Printf("Host: %s", o.Broker)

	if o.TLS {
		opts.SetTLSConfig(NewTLSConfig())
	}

	if o.SetWill {
		stateReq := &StateRequest{Timestamp: 0, State: StateOFF}
		msg, err := json.Marshal(stateReq)
		if err != nil {
			log.Fatal(err)
		}

		stateTopic := GetTopicString(o.Serial, StateTopic)
		opts.SetWill(stateTopic, string(msg), o.QoS, o.Retained)
	}

	return Client{o: o, mc: mqtt.NewClient(opts)}
}
