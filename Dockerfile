FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/portfolio/main.go

RUN mkdir /app/credentials

FROM alpine:latest

COPY --from=builder /app/api /

COPY --from=builder /app/credentials /credentials

CMD ["./api"]