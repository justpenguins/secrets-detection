FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .


RUN go build -o secrets main.go
# FROM gcr.io/distroless/base-debian11


WORKDIR /app
COPY --from=builder /app/secrets /app/secrets


CMD ["/app/secrets"]