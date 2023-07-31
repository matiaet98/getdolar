FROM golang:1.20-alpine3.18 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o getdolar main.go

FROM alpine:3.18
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/getdolar /usr/local/bin/getdolar

CMD ["/usr/local/bin/getdolar"]
