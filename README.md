# count-server

## Usage
Run the count-server:
```
❯ go run server/main.go 
```

In another terminal, traverse to root directory of this repo and then hit the server by running:
```
❯ buf curl \
--schema proto \
http://localhost:8080/count.v1.CountService/Count
```

## Useful Commands
Verify that the proto is valid:
```
❯ buf build proto
❯ echo $?
```
The output should be `0`

Generate code:
```
❯ buf generate proto
```

Lint the API:
```
❯ buf lint proto
```

Update buf dependencies to the latest versions:
```
❯ cd proto
❯ buf mod update
```

Push new versions of the API:
```
❯ buf push proto
```
