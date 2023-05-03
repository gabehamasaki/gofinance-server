FROM golang:1.20-alpine3.16 as base
RUN apk update
WORKDIR /src/server
COPY go.mod go.sum ./
COPY . . 
RUN go build -o gofinance ./cmd/server

FROM alpine:3.16 as binary
COPY --from=base /src/server/gofinance .
EXPOSE 3000
CMD ["./gofinance"]