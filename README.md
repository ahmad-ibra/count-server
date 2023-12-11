# count-server

## Usage
Run the count-server container via docker:
```
❯ docker build -t count-server .
❯ docker network create count-network
❯ docker run -p 8080:8080 --network=count-network --name count-server -d count-server
```

In another terminal, traverse to root directory of this repo and then hit the server by running:
```
❯ buf curl \
--schema proto \
http://0.0.0.0:8080/count.v1.CountService/Count
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

Push new version of the image:
```
❯ docker build -t ahmadibraspectrocloud/count-server:latest .
❯ docker push ahmadibraspectrocloud/count-server:latest
```
