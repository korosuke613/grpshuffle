# Example of grpshuffle for Node.js (client)
Start the client-server using the express.

## Requirement

### Build grpshuffle server
```
cd ../../
make server
```

## Getting started

### Launch grpshuffle server
```console
../../server
```

### Launch groshuffle client-server

```console
npx ts-node ./index.ts
```

### Access client-server
```console
❯ curl "localhost:8080/shuffle?partition=2&targets=a,11,15"
{"result":{"combinations":[{"targets":["15","a"]},{"targets":["11"]}]}}
```
