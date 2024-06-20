# go-log

## Description

The goal of this project is to provide a simple and easy to use binary which receives a command as an argument and logs the output of the command. The output can be logged to a file or to the standard output. The loglevel can be set to different levels according to the [`zerolog package`](https://pkg.go.dev/github.com/rs/zerolog).

## Why this project?

I created this project because I needed a simple way to log the output of a command in order to monitor the execution of a script/program/command. I faced some issues with Subprocesses in Python (timeout, graceful shutdown, etc) and I decided to create a simple binary in Go to solve this problem by using the [`os/exec package`](https://pkg.go.dev/os/exec) and the [`context`](https://pkg.go.dev/context) package.

## How to build it?

The following command will build several binaries for different operating systems and architectures and place them in the [`build`](./build/) directory:

```bash
make build
```

A simple `make` command will build only one binary with the machine configuration:

```bash
make
```

## Requirements

- Built through the **Go**'s version **1.22.4**
- `make` installed

## Dependencies

- [`github.com/rs/zerolog`](https://pkg.go.dev/github.com/rs/zerolog): For logging the output of the command
- [`github.com/spf13/pflag`](https://pkg.go.dev/github.com/spf13/pflag): For parsing the command line arguments

## How to use it?

The **only one required** argument is the `--command` flag, the other flags are **optional** and have default values which can be observed by running the binary with the `--help` flag.

The following command will log the output of the command `ls -la` to the file `output.log` with the loglevel set to `info`:

```bash
./build/go-log --command="ls -la" --loglevel "info" --output "output.log"
```

You can also provide a timeout value to the command execution defined in `seconds`, the following command will log the output of the command `sleep 5` to the standard output with the a timeout of 1 second, it will produce an error message because the command will take more than 1 second to execute:

```bash
./build/go-log --command="sleep 5" --timeout=1
```
