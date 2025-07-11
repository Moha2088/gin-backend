package config

import "go.uber.org/zap"

func SetupLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	if err != nil {
		panic(err)
	}
	return logger
}
