package cli

import (
	"errors"

	"go-log/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func ParseFlag() zerolog.Logger {
	var logger zerolog.Logger
	var version bool
	var help bool

	flag.BoolVarP(&help, "help", "h", false, "The help flag")
	flag.BoolVarP(&version, "version", "v", false, "The version flag")
	flag.StringVarP(&cmd.Output, "output", "o", "", "The destination file output")
	flag.StringVarP(&cmd.Command, "command", "c", "", "String which contains command")
	flag.StringVarP(&cmd.LogLevel, "loglevel", "l", "info", "String which defines the zerolog/loglevel")
	flag.IntVarP(&cmd.TimeOut, "timeout", "t", cmd.TimeOut, "Integer which defines the timeout accepted duration")
	flag.Usage = printHelp
	flag.Parse()

	if help {
		handleHelp(0)
	}
	if version {
		handleVersion(0)
	}
	if cmd.Command == "" {
		log.Error().Err(errors.New("command flag is empty")).Msg("you didn't provide a value for the command flag")
		handleHelp(1)
	}
	if cmd.TimeOut <= 0 {
		log.Fatal().Err(errors.New("timeout flag unsupported value: timeout <= 0")).Msg("please provide a timeout's value > 0")
	}
	logger = handleLog()
	return logger
}
