# Log Level

The log level is a way to control the verbosity of the logs. The log level can be set to one of the following values:

- `trace`   The **most** verbose log level
- `debug`   The **second most** verbose log level
- `info`    The **default** log level
- `warn`    The **second least** verbose log level
- `error`   The **least** verbose log level
- `fatal`   The log level that **will stop the program execution** in case of an error.
- `panic`   The log level that **will stop the program execution** and **print a stack trace** in case of an error.

## Log Level Example

The following command will log the output of the command `echo "hello world"` to the file `output.log` with the loglevel set to `info`:

```bash
./build/go-log --command="echo \"hello world\"" --loglevel "info" --output "output.log"
```

```bash
cat output.log
```

```json
{"level":"info","pid":177057,"exitCode":"0","time":"2024-06-28T19:02:32+02:00","message":"hello world"}
```

## Log Level Error

The following command will log the output of the command `ls -la /nonexistent` to the standard output with the loglevel set to `info`, the command will produce an error message because the directory `/nonexistent` doesn't exist:

```bash
./build/go-log --command="ls -la /nonexistent" --loglevel "info"
```

```json
{"level":"fatal","pid":177670,"error":"error waiting for command: exit status 2","time":"2024-06-28T19:03:20+02:00","message":"an error has occurred"}
```

## Log Level Trace

The following command will log the output of the command `echo "hello world"` to the standard output with the loglevel set to `trace`:

```bash
./build/go-log --command="echo \"hello world\"" --loglevel "trace"
```

```json
{"level":"trace","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"command: echo \"hello world\""}
{"level":"trace","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"loglevel: trace"}
{"level":"trace","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"the program is starting"}
{"level":"debug","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"starting running the command: echo \"hello world\""}
{"level":"debug","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"command finished in 895.066µs"}
{"level":"info","pid":178323,"exitCode":"0","time":"2024-06-28T19:10:07+02:00","message":"hello world"}
{"level":"trace","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"the command has finished"}
{"level":"trace","pid":178323,"time":"2024-06-28T19:10:07+02:00","message":"the program is ending"}
```

## Log Level Debug

The following command will log the output of the command `echo "hello world"` to the standard output with the loglevel set to `debug`:

```bash
./build/go-log --command="echo \"hello world\"" --loglevel "debug"
```

```json
{"level":"debug","pid":178820,"time":"2024-06-28T19:11:07+02:00","message":"starting running the command: echo \"hello world\""}
{"level":"debug","pid":178820,"time":"2024-06-28T19:11:07+02:00","message":"command finished in 2.959931ms"}
{"level":"info","pid":178820,"exitCode":"0","time":"2024-06-28T19:11:07+02:00","message":"hello world"}
```
