package logging

import (
	"encoding/json"
	"go.uber.org/zap"
)

var LOGGER *zap.Logger

func InitLogging() {
	rawJsonConfig := []byte(`
	{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		  "encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		  }
	}`,
	)

	var config zap.Config
	if err := json.Unmarshal(rawJsonConfig, &config); err != nil {
		panic(err)
	}

	LOGGER = zap.Must(config.Build())

	LOGGER.Info("Starting up")
}
