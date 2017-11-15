package impl

import logging "github.com/reflect/reflect-logging"

var (
	logger = logging.Builder().At("errawr-go", "impl")
)

func log() logging.Logger {
	return logger.Build()
}
