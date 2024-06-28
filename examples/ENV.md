# Env

The environment variables file can be used to set the environment variables for the binary. They can be used in the command to be executed.

## Success Example

```bash
cat .env-example
VALUE=HELLO WORLD
```

```bash
 ./build/go-log --env=.env-example --command="sh -c 'echo \$VALUE'" --loglevel=trace
```

```json
{"level":"debug","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"command: sh -c 'echo $VALUE'"}
{"level":"debug","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"loglevel: trace"}
{"level":"debug","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"env: .env-example"}
{"level":"trace","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"the program is starting"}
{"level":"debug","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"starting running the command: sh -c 'echo $VALUE'"}
{"level":"debug","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"command finished in 1.78415ms"}
{"level":"info","pid":193415,"exitCode":"0","time":"2024-06-28T20:01:52+02:00","message":"HELLO WORLD"}
{"level":"trace","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"the command has finished"}
{"level":"trace","pid":193415,"time":"2024-06-28T20:01:52+02:00","message":"the program is ending"}
```

## Error Example

```bash
cat .env-undefined
```

```output
cat: .env-undefined: No such file or directory
```

```bash
 ./build/go-log --env=.env-undefined --command="sh -c 'echo \$VALUE'" --loglevel=trace
```

```json
{"level":"fatal","error":"open .env-undefined: no such file or directory","time":"2024-06-28T20:05:40+02:00","message":"error setting the environment variables"}
```
