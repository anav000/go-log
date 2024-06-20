package tools

import (
	"os"

	"go-log/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CreateLogger() zerolog.Logger {
	var logger zerolog.Logger

	if cmd.Output != "" {
		file, err := os.OpenFile(cmd.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal().Err(err).Msgf("error when trying to open the path: %s", cmd.Output)
		}
		logger = zerolog.New(file).
			With().
			Timestamp().
			Int("pid", os.Getpid()).
			Logger()
	} else {
		logger = zerolog.New(os.Stderr).
			With().
			Timestamp().
			Int("pid", os.Getpid()).
			Logger()
	}
	logger.Trace().Msgf("command: %s", cmd.Command)
	logger.Trace().Msgf("loglevel: %s", log.Logger.GetLevel())
	if cmd.Output != "" {
		logger.Trace().Msgf("output: %s", cmd.Output)
	}
	return logger
}
