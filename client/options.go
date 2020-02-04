package main

// Opts for application
type Opts struct {
	Broker   string
	ID       string
	Serial   string
	Mode     Mode
	QoS      byte
	Retained bool
	SetWill  bool
	TLS      bool
}

// NewOpts return opts struct with default values
func NewOpts() Opts {
	return Opts{
		Broker:   getenv("MQTT_TEST_CLIENT_HOST", "ssl://mqtt:8883"),
		ID:       getenv("MQTT_TEST_CLIENT_ID", "test-client"),
		Serial:   "1001",
		Mode:     ModeAuto,
		QoS:      2,
		Retained: true,
		SetWill:  false,
		TLS:      false,
	}
}
