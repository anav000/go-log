package cli

import (
	"errors"
	"os"
	"slices"

	"go-log/cmd"
	"go-log/tools"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func handleHelp(rc int) {
	printHelp()
	os.Exit(rc)
}

func handleVersion(rc int) {
	printVersion()
	os.Exit(rc)
}

func handleLog() zerolog.Logger {
	var logger zerolog.Logger
	var level zerolog.Level
	var err error

	defaultErr := errors.New("the loglevel given by the user isn't supported")
	if cmd.LogLevel == "" {
		log.Error().Err(defaultErr).Msgf("Available log levels: %v", getLogLevel())
		handleHelp(1)
	}
	if !slices.Contains(getLogLevel(), cmd.LogLevel) {
		log.Fatal().Err(defaultErr).Msgf("Available log levels: %v", getLogLevel())
	}
	if level, err = zerolog.ParseLevel(cmd.LogLevel); err != nil {
		log.Fatal().Err(err).Msgf("Available log levels: %v", getLogLevel())
	}
	zerolog.SetGlobalLevel(level)
	logger = tools.CreateLogger()
	return logger
}
