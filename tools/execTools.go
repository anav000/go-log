package tools

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func CreatePipe(cmd *exec.Cmd, pipeType string) (io.ReadCloser, error) {
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

func ReadStandardBuffer(stdChan chan string, stdPipe io.ReadCloser) {
	var stdBuffer bytes.Buffer

	io.Copy(&stdBuffer, stdPipe)
	stdChan <- stdBuffer.String()
	close(stdChan)
}
