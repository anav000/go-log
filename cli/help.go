package cli

import (
	"fmt"

	"go-log/cmd"
)

func getLogLevel() []string {
	return []string{"debug", "error", "fatal", "info", "log", "panic", "trace", "warn"}
}

func printHelp() {
	fmt.Printf("%s\n\n", cmd.Banner)
	fmt.Printf("USAGE\tgo-log --command command [--output output] [--loglevel level] [--timeout duration]\n\n")
	fmt.Printf("\t-v, --version\tDisplay the project's version and exit\n")
	fmt.Printf("\t-h, --help\tDisplay the project's usage and exit\n")
	fmt.Printf("\t-o, --output\tDefine the file output\n")
	fmt.Printf("\t-c, --command\tDefine the command to be run\n")
	fmt.Printf("\t-t, --timeout\tDefine the timeout allowed duration (default: %d)\n", cmd.TimeOut)
	fmt.Printf("\t-l, --loglevel\tDefine the log level (default: %s)\n", cmd.LogLevel)
	fmt.Printf("\navailable levels: %v\n", getLogLevel())
}
