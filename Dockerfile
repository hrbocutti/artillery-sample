FROM golang:1.22.3 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build cmd/app/main.go

FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /root/

COPY --from=builder /app/main .
RUN chmod +x ./main

EXPOSE 8081
CMD ["./main"]
