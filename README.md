# Golang hits

this app by default listen on port 80, you can override the the value with set environment variable `HTTP_PORT`

example output:
```
User address: <client remote address>
Hits: <client hits increment>
```

## Build docker image

Using docker build
```
docker build -t golang-hits .
```

## Run Application

golang-hits app will run on backend, frontend by traefik for reverse proxy and loadbalancer

```
docker-compose up -d --build
```

access traefik dashboard on port localhost:8080

access app using curl
```
curl localhost
```

## Scale App

```
docker-compose up -d --scale golang-hits=<number>
```

## Shutdown Application

```
docker-compose down
```