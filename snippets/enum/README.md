
### tool: stringer

install `golang.org/x/tools/cmd/stringer`

and running this command:

```bash
stringer -type=Pill
```

in the same directory will create the file `pill_string.go`, this file will contain a definition of:

```go
func (Pill) String() string
```
