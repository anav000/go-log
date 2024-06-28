package cli

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"go-log/cmd"
	"go-log/tools"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// handleHelp prints the help message and exits with the specified status code.
func handleHelp(rc int) {
	printHelp()
	os.Exit(rc)
}

// handleVersion prints the version message and exits with the specified status code.
func handleVersion(rc int) {
	printVersion()
	os.Exit(rc)
}

// handleLog handles the log level given by the user.
// The function returns a logger with the specified log level.
// If the log level given by the user isn't supported, the function prints an error message and exits with status code 1.
func handleLog(envFile string) zerolog.Logger {
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
	logger = tools.CreateLogger(envFile)
	return logger
}

// handleEnv handles the environment variables file given by the user.
// The function reads the file and sets the environment variables.
// If an error occurs, the function prints an error message and exits with status code 1.
func handleEnv(env string) {
	if env == "" {
		return
	}
	err := setEnv(env)
	if err != nil {
		log.Fatal().Err(err).Msg("error setting the environment variables")
	}
}

// setEnv reads the environment variables file and sets the environment variables.
// The function returns an error if one occurred.
func setEnv(envFile string) error {
	env, err := godotenv.Read(envFile)
	if err != nil {
		return err
	}
	for key, value := range env {
		formattedEnv := fmt.Sprintf("%s=%s", key, value)
		cmd.Env = append(cmd.Env, formattedEnv)
	}
	return nil
}
