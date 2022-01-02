# grpshuffle
A gRPC server and client that randomly assigns targets.

[![CI](https://github.com/korosuke613/grpshuffle/actions/workflows/ci.yaml/badge.svg)](https://github.com/korosuke613/grpshuffle/actions/workflows/ci.yaml)
[![](https://img.shields.io/badge/protocol-doc-blue)](./doc/grpshuffle.md)
[![release go](https://img.shields.io/github/v/release/korosuke613/grpshuffle?color=007d9c&logo=go)](https://github.com/korosuke613/grpshuffle/releases)
[![document go](https://pkg.go.dev/badge/github.com/korosuke613/grpshuffle/go/grpshuffle)](https://pkg.go.dev/github.com/korosuke613/grpshuffle/go)
[![release npm](https://img.shields.io/npm/v/grpshuffle-js?color=EA2039&logo=npm&label=release)](https://www.npmjs.com/package/grpshuffle-js)

## Getting started
### 1. Build server & client
```
make server
make client
```

### 2. Launch server
```
./server
```

### 3. Execute client
```
./client 127.0.0.1:13333 shuffle 2 a b c d e
```

result: 
```json
[
  {
    "targets": [
      "a",
      "c",
      "b"
    ]
  },
  {
    "targets": [
      "e",
      "d"
    ]
  }
]
```

## Install

### Server
```
go install github.com/korosuke613/grpshuffle/go/grpshuffle-server@latest
```

### Client
```
go install github.com/korosuke613/grpshuffle/go/grpshuffle-client@latest
```

## Usage

### client
```
usage: client HOST:PORT METHOD PARTITION TARGET_1 TARGET_2 ... TARGET_N
METHOD: shuffle, health
```

## Build
```
make
```

## Release

1. Create git tag (`git tag v0.0.x`)
2. Push git tag (`git push origin v0.0.x`)
3. Auto publish by GitHub Actions (https://github.com/korosuke613/grpshuffle/actions/workflows/release.yaml)
