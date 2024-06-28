# Timeout Example

This file contains a list of examples of how to use the `go-log` binary with the `--timeout` flag.

## Timeout Error

The following command will log the output of the command `sleep 5` to the standard output with a `timeout of 1 second`, it will produce an error message because the command will take `more than 1 second` to execute:

```bash
./build/go-log --command="sleep 5" --timeout=1
```

```json
{"level":"fatal","pid":165284,"error":"context deadline exceeded","time":"2024-06-28T18:39:45+02:00","message":"the command has timed out"}
```

## Timeout Success

The following command will log the output of the command `sleep 1` to the standard output with a `timeout of 2 seconds`, the command won't produce an error message because the command will take `less than 2 seconds` to execute:

```bash
./build/go-log --command="sleep 1" --timeout=2
```
