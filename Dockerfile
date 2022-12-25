FROM golang:1.17.1-alpine3.14
# Enable go modules
ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app


# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

RUN go mod download



WORKDIR /golang-app
COPY . .
# Build the application.
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o golang-app .


# Finally our multi-stage to build a small image
# Start a new stage from scratch
FROM alpine:3.14

WORKDIR /app    

# Copy the Pre-built binary file
COPY --from=0 /golang-app .


EXPOSE 3000

# Run executable
CMD ["./golang-app"]
