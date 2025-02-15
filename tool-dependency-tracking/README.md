See also [Tool dependencies](https://go.dev/doc/modules/managing-dependencies#tools)

# Track a tool dependency

```shell
go get -tool golang.org/x/tools/cmd/stringer
```

The `go tool` command can then run these tools in addition to tools shipped with the Go distribution.

```shell
go tool stringer $args
```
