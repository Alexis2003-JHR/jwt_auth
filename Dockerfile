FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd/

RUN go mod tidy
RUN go build -o /app/jwt-auth

# Stage 2: Run the application using a smaller base image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/jwt-auth /app/jwt-auth

RUN chmod +x /app/jwt-auth

COPY private.key /app/private.key
COPY public.key /app/public.key

CMD ["./jwt-auth"]
