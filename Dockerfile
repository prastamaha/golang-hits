FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .

RUN go build -o /golang-hits

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /golang-hits .

EXPOSE 80

CMD ["/app/golang-hits"]

