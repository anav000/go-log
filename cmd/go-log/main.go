package main

import (
	"context"
	"time"

	"go-log/cli"
	"go-log/cmd"
	"go-log/tools"
)

// main is the entry point of the program.
func main() {
	running := true
	logger := cli.ParseFlag()
	done := make(chan struct{})
	errChan := make(chan error)
	timeout := time.Duration(cmd.TimeOut) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	defer close(errChan)

	logger.Trace().Msg("the program is starting")
	go tools.ExecCommand(ctx, logger, errChan, done)

	for running {
		select {
		case <-done:
			logger.Trace().Msg("the command has finished")
			running = false
		case err := <-errChan:
			logger.Fatal().Err(err).Msg("an error has occurred")
		case <-ctx.Done():
			logger.Fatal().Err(ctx.Err()).Msg("the command has timed out")
		}
	}
	logger.Trace().Msg("the program is ending")
}
