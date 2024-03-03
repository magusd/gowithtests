
# Gowithtests

# Goland refactor 

- CMD+OPTION+M -> extract method
- CMD+OPTION+R -> refactor menu
- CMD+OPTION+B -> go to declaration


# Go Docs

to serve the docs with the current project's examples and documentation
```bash
godoc -http=:6060
```

to get the docs about a function
```bash
go doc fmt.Println
```

# Benchmark

```bash
go test -v -bench=. ./...
```

# Coverage

```bash
go test -cover ./...
```


# Build without debug and symbol table

```bash
go build -ldflags "-w -s"
```

# Build for multiple OS and Arch

```bash
GOOS="linux" GOARCH="amd64" go build hello.go
```