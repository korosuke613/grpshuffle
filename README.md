# grpshuffle
A gRPC server and client that randomly assigns targets.

[![CI](https://github.com/korosuke613/grpshuffle/actions/workflows/ci.yaml/badge.svg)](https://github.com/korosuke613/grpshuffle/actions/workflows/ci.yaml)
[![](https://img.shields.io/badge/protocol-doc-blue)](./doc/grpshuffle.md)

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
./client 127.0.0.1:13333 2 a b c d e
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

## Usage

### client
```
usage: client HOST:PORT PARTITION TARGET_1 TARGET_2 ... TARGET_N
```

## Build
```
make
```

## Release

1. Create git tag (`git tag v0.0.x`)
2. Push git tag (`git push origin v0.0.x`)
3. Auto publish by GitHub Actions (https://github.com/korosuke613/grpshuffle/actions/workflows/release.yaml)
