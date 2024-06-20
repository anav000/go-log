package main

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"go-log/cli"
	"go-log/cmd"
	"go-log/tools"

	"github.com/rs/zerolog"
)

func execCommand(ctx context.Context, logger zerolog.Logger, errChan chan error, done chan struct{}) int {
	execCmd := exec.CommandContext(ctx, "sh", "-c", cmd.Command)
	logger.Debug().Msgf("starting running the command: %s", cmd.Command)
	stdoutPipe, errStdout := tools.CreatePipe(execCmd, "stdout")
	stderrPipe, errStderr := tools.CreatePipe(execCmd, "stderr")
	stdoutChan := make(chan string)
	stderrChan := make(chan string)

	defer stdoutPipe.Close()
	defer stderrPipe.Close()

	if errStdout != nil {
		errChan <- fmt.Errorf("error creating stdout pipe: %v", errStdout)
		return 1
	}
	if errStderr != nil {
		errChan <- fmt.Errorf("error creating stdout pipe: %v", errStderr)
		return 1
	}

	timeDuration := time.Now()
	if err := execCmd.Start(); err != nil {
		errChan <- fmt.Errorf("error starting command: %v", err)
		return 1
	}

	go tools.ReadStandardBuffer(stdoutChan, stdoutPipe)
	go tools.ReadStandardBuffer(stderrChan, stderrPipe)

	if err := execCmd.Wait(); err != nil {
		errChan <- fmt.Errorf("error waiting for command: %v", err)
		return execCmd.ProcessState.ExitCode()
	}

	stdoutStr := <-stdoutChan
	stderrStr := <-stderrChan
	exitCode := execCmd.ProcessState.ExitCode()
	logger.Debug().Msgf("the command has been executed in %s", time.Since(timeDuration))
	if stdoutStr != "" {
		logger.Info().Str("exitCode", strconv.Itoa(exitCode)).Msg(strings.TrimSuffix(stdoutStr, "\n"))
	}
	if stderrStr != "" {
		logger.Info().Str("exitCode", strconv.Itoa(exitCode)).Msg(strings.TrimSuffix(stderrStr, "\n"))
	}
	close(done)
	return exitCode
}

func main() {
	running := true
	logger := cli.ParseFlag()
	done := make(chan struct{})
	errChan := make(chan error)
	timeout := time.Duration(cmd.TimeOut) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Trace().Msg("the program is starting")
	go execCommand(ctx, logger, errChan, done)

	for running {
		select {
		case <-done:
			logger.Trace().Msg("the go routine has finished")
			running = false
		case err := <-errChan:
			logger.Fatal().Err(err).Msg("an error has occured during the execution of the command")
		case <-ctx.Done():
			logger.Fatal().Err(ctx.Err()).Msg("the routine badly ended")
		}
	}
	logger.Trace().Msg("the program is ending")
}
