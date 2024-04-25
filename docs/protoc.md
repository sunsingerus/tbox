# Required components
1. `go` version 1.17 at least
1. `protoc`
1. `protoc-gen-go`
1. `protoc-gen-go-grpc`
1. `protoc-gen-grpc-web`
1. `protoc-gen-doc`

## How to install `protoc`

- Download latest `protoc` release from [here](https://github.com/protocolbuffers/protobuf/releases)
- We'll have something like `protoc-3.19.4-linux-x86_64.zip` with the following structure:
```text
    bin
        protoc
    include
        google
            protobuf
                ... many files here ...
```
- Place `bin` content into `$PATH`-searchable - `bin`
- Place `include` near `bin`, so we'll have something like the following:
```text
    bin
        ... $PATH-searchable bin folder ...
        ... you may have your old bin files ...
        protoc
    include
        google
            protobuf
                ... many files here ...
``` 

## How to install `protoc-gen-go`

Outside any go-related source code folder (say in `/` or `/home`), in order for `go` not to use any `go.mod` which it can find run:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
OR if you are brave
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
which is described here:
```
https://pkg.go.dev/google.golang.org/protobuf
```
and which lives in this repo
```
https://github.com/protocolbuffers/protobuf-go
```
and is located here
```
https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go
```
You should have `protoc-gen-go` executable, which has to be available in `$PATH` for being called
Such as
```bash
mv ~/gopath/bin/protoc-gen-go ~/bin/
```

## How to install `protoc-gen-go-grpc`

Outside any go-related source code folder (say in `/` or `/home`), in order for `go` not to use any `go.mod` which it can find run:
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
OR if you are brave
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
which is described here:
```
https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
which lives in this repo
```
https://github.com/grpc/grpc-go
```
and is located here
```
https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc
```
You should have `protoc-gen-go-grpc` executable, which has to be available in `$PATH` for being called
Such as
```bash
mv ~/gopath/bin/protoc-gen-go-grpc ~/bin/protoc-gen-go-grpc
```

## How to install `protoc-gen-grpc-web`

Get `grpc-web` generator as
```bash
https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64
mv protoc-gen-grpc-web-1.2.1-linux-x86_64 protoc-gen-grpc-web
chmod a+x protoc-gen-grpc-web
```
Move `grpc-web` generator into `$PATH`-listed folder as 
```bash
mv protoc-gen-grpc-web ~/bin/
```

## How to install `protoc-gen-doc`

Outside any go-related source code folder (say in `/` or `/home`), in order for `go` not to use any `go.mod` which it can find run:
```bash
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.0
OR if you are brave
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
```
which lives in this repo
```
https://github.com/pseudomuto/protoc-gen-doc
```
and is located here
```
https://github.com/pseudomuto/protoc-gen-doc/tree/master/cmd/protoc-gen-doc
```
You should have `protoc-gen-doc` executable, which has to be available in `$PATH` for being called
Such as
```bash
mv ~/gopath/bin/protoc-gen-doc ~/bin/protoc-gen-doc
```

# Check all components are in place

```bash
which protoc protoc-gen-go protoc-gen-go-grpc protoc-gen-grpc-web protoc-gen-doc
```
should produce something like
```text
~/bin/protoc
~/bin/protoc-gen-go
~/bin/protoc-gen-go-grpc
~/bin/protoc-gen-grpc-web
~/bin/protoc-gen-doc
```
