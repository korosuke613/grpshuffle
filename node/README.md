# grpshuffle-client for Node.js

## grpshuffle-client-server
Start the client-server using the express.

### Pre-Requirement

#### Build grpshuffle server
```
cd ../
make server
```

#### Launch grpshuffle server
```console
../server
```


### Getting started

#### Launch groshuffle client-server

```console
npx -p express-generator -c "grpshuffle-client-server"
```

#### Access client-server
```console
❯ curl "localhost:8080/shuffle?divide=2&targets=a,11,15"
{"result":{"combinations":[{"targets":["15","a"]},{"targets":["11"]}]}}
```

## Development

### Setup

#### npm install
```console
npm i
```

#### Generate library
```console
make
```

### Publish
See [release-node.yaml](/.github/workflows/release-node.yaml).
