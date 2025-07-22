# Dockerfile
FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o otp-auth ./main.go

EXPOSE 8080

CMD ["./otp-auth"]
