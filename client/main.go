package main

import (
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

func main() {
	o := NewOpts()

	modePrompt := promptui.Select{
		Label: "Select action:",
		Items: []string{ModeAction},
	}

	_, action, err := modePrompt.Run()

	if err != nil {
		log.Errorf("Mode prompt failed %v\n", err)
		return
	}

	switch action {
	case ModeAction:
		o.PromptMode()
		_, err := o.PromptModifyDefaults()
		if err != nil {
			log.Fatal(err)
		}
		c := NewClient(o)
		c.PublishMode()
		break
	}

}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// GetTopicString return MQTT topic for a given serial and subtopic
func GetTopicString(serial string, subTopic Topic) string {
	topic := strings.Join([]string{
		"/v1.0",
		serial,
		string(subTopic),
		"format",
		"json"}, "/")
	return topic
}
