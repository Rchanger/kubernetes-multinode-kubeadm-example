
# FROM alpine  
FROM golang:1.13.5-alpine3.11 as builder

# Add Maintainer Info
LABEL maintainer="Suvarna Rokade"

# Install git
RUN apk --update add git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine  

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder go/src/app/server .

# Expose port 8090 to the outside world
EXPOSE 8090

RUN chmod a+x server

# Command to run the executable
CMD ["./server"]

