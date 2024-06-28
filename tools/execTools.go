package tools

import (
	"bytes"
	"context"
	"fmt"
	"go-log/cmd"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// CreatePipe creates a pipe for the specified command.
// The pipeType parameter specifies the type of pipe to create (stdout or stderr).
// The function returns the pipe and an error if one occurred.
func createPipe(cmd *exec.Cmd, pipeType string) (io.ReadCloser, error) {
	var pipe io.ReadCloser
	var err error

	switch pipeType {
	case "stdout":
		pipe, err = cmd.StdoutPipe()
	case "stderr":
		pipe, err = cmd.StderrPipe()
	default:
		return nil, fmt.Errorf("unknown pipe type: %s", pipeType)
	}

	if err != nil {
		return nil, fmt.Errorf("error creating %s pipe: %w", pipeType, err)
	}
	return pipe, nil
}

// ReadStandardBuffer reads the standard buffer and sends it to the specified channel.
// The function closes the channel when the buffer is read.
func readStandardBuffer(stdChan chan string, stdPipe io.ReadCloser) {
	var stdBuffer bytes.Buffer

	io.Copy(&stdBuffer, stdPipe)
	stdChan <- stdBuffer.String()
	close(stdChan)
}

// execCommand executes the command specified in the command line.
// The function returns the exit code of the command.
// If an error occurs, the function sends the error to the specified channel.
// The function also sends a signal to the done channel when the command has finished.
func ExecCommand(ctx context.Context, logger zerolog.Logger, errChan chan error, done chan struct{}) int {
	execCmd := exec.CommandContext(ctx, "sh", "-c", cmd.Command)
	execCmd.Env = append(execCmd.Env, cmd.Env...)
	stdoutPipe, errStdout := createPipe(execCmd, "stdout")
	stderrPipe, errStderr := createPipe(execCmd, "stderr")
	stdoutChan := make(chan string)
	stderrChan := make(chan string)

	defer stdoutPipe.Close()
	defer stderrPipe.Close()
	defer close(done)

	logger.Debug().Msgf("starting running the command: %s", cmd.Command)
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

	go readStandardBuffer(stdoutChan, stdoutPipe)
	go readStandardBuffer(stderrChan, stderrPipe)

	if err := execCmd.Wait(); err != nil {
		errChan <- fmt.Errorf("error waiting for command: %v", err)
		return execCmd.ProcessState.ExitCode()
	}

	stdoutStr := <-stdoutChan
	stderrStr := <-stderrChan
	exitCode := execCmd.ProcessState.ExitCode()
	logger.Debug().Msgf("command finished in %s", time.Since(timeDuration))
	if stdoutStr != "" {
		logger.Info().Str("exitCode", strconv.Itoa(exitCode)).Msg(strings.TrimSuffix(stdoutStr, "\n"))
	}
	if stderrStr != "" {
		logger.Info().Str("exitCode", strconv.Itoa(exitCode)).Msg(strings.TrimSuffix(stderrStr, "\n"))
	}
	return exitCode
}
