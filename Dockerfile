# Base Image
FROM golang:1.18.3-alpine3.16 as base

# Working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o server ./src/.

# Create master image
FROM alpine AS master

# Working directory
WORKDIR /app

# Copy execute file
COPY --from=base /app/server ./

# add iwatch
RUN apk add --no-cache iwatch=0.2.2-r0

# Set ENV to production
ENV GO_ENV production

# Expose port 3001
EXPOSE 3001

# Run the application
CMD ["./server"]