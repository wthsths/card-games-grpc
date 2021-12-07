FROM golang:1.17 as builder

WORKDIR /usr/local/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/card-games-grpc cmd/server/main.go 

FROM alpine:3.15
WORKDIR /root

COPY --from=builder /usr/local/bin/card-games-grpc ./

CMD ["./card-games-grpc"]