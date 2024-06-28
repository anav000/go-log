# Output Example

This file contains a list of examples of how to use the `go-log` binary with the `--output` flag. If the file is not provided, the output will be logged to the standard output. In case the file doesn't exist, it will be created, else the content will be appended to the file.

## Output to File

The following command will log the output of the command `echo "hello world"` to the file `output.log` with the loglevel set to `info`:

```bash
./build/go-log --command="echo \"hello world\"" --loglevel "info" --output "output.log"
```

```bash
cat output.log
```

```json
{"level":"info","pid":172092,"exitCode":"0","time":"2024-06-28T18:49:54+02:00","message":"hello world"}
```

## Output to Standard Output

The following command will log the output of the command `echo "hello world"` to the standard output with the loglevel set to `info`:

```bash
./build/go-log --command="echo \"hello world\"" --loglevel "info"
```

```json
{"level":"info","pid":173397,"exitCode":"0","time":"2024-06-28T18:54:28+02:00","message":"hello world"}
```

## Output to Standard Output with Error

The following command will log the output of the command `ls -la /nonexistent` to the standard output with the loglevel set to `info`, the command will produce an error message because the directory `/nonexistent` doesn't exist:

```bash
./build/go-log --command="ls -la /nonexistent" --loglevel "info"
```

```json
{"level":"fatal","pid":173829,"error":"error waiting for command: exit status 2","time":"2024-06-28T18:55:20+02:00","message":"an error has occurred"}
```
