# Use a newer Go image that matches your required Go version
FROM golang:1.23.2 AS builder

# Continue with the rest of your Dockerfile commands
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN go build -o app

# Now, you can use a smaller image to run the built binary
FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
