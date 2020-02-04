package main

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

// PublishMode handle mode publish to broker
func (c *Client) PublishMode() {
	if token := c.mc.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Connection Failed: %e", token.Error())
	}

	modeReq := &ModeRequest{
		Mode:      &c.o.Mode,
		Timestamp: time.Now().Unix()}

	msg, err := json.Marshal(modeReq)
	if err != nil {
		log.Fatal(err)
	}

	// Send Mode
	resource := GetTopicString(c.o.Serial, ModeTopic)
	token := c.mc.Publish(resource, c.o.QoS, c.o.Retained, msg)

	if token.WaitTimeout(time.Second*4) == false {
		log.Fatalf("Publish timeout to broker: %s", c.o.Broker)
	}
	if token.Error() != nil {
		log.Fatalf("Publish failed: %e", token.Error())
	}
}
