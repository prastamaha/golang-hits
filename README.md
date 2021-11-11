# Golang hits

this app by default listen on port 80, you can override the the value with set environment variable `HTTP_PORT`

example output:
```
User address: <client remote address>
Hits: <client hits increment>
```

## Build docker image

```
docker build -t golang-hits .
```

