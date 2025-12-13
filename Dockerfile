FROM golang:alpine AS builder

# Stage 1 : dependencies download
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /tq

FROM alpine:latest
# Stage 2 : Simply deploy the built bin
WORKDIR /app
COPY handlers /app/handlers
COPY templates /app/templates
COPY --from=builder /tq .

EXPOSE 8080

CMD ["/app/tq"]