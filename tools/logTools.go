package tools

import (
	"os"

	"go-log/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// CreateLogger creates a logger with the specified output file.
// If the output file is not specified, the logger will write to stderr.
// The logger will also log the command, loglevel and output file if specified.
func CreateLogger(envFile string) zerolog.Logger {
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
	logger.Debug().Msgf("command: %s", cmd.Command)
	logger.Debug().Msgf("loglevel: %s", log.Logger.GetLevel())
	if cmd.Output != "" {
		logger.Debug().Msgf("output: %s", cmd.Output)
	}
	if envFile != "" {
		logger.Debug().Msgf("env: %s", envFile)
	}
	return logger
}
