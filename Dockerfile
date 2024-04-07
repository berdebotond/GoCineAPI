# Step 1: Use an official Go runtime as a parent image
FROM golang:1.22 as builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the go.mod and go.sum to download dependencies
# Copying just these files first improves the Docker cache layers, as dependencies are less likely to change than our application code
COPY go.mod go.sum ./

# Step 4: Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Step 5: Copy the source code into the container
COPY . .

# Step 6: Build the Go app for a linux OS with disabled CGO. This creates a smaller and more secure binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main

# Step 7: Use a lightweight, secure base image (e.g., Alpine) for the final image
FROM alpine:latest  

# Step 8: Install CA certificates for secure communication
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Step 9: Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Step 10: Copy any other necessary files or directories (like your config directory)
COPY --from=builder /app/config ./config

# Step 11: Expose port 8123 to the outside world
EXPOSE 8123

# Step 12: Run the binary program
CMD ["./main"]
