# FROM alpine  
FROM golang:1.13.5-alpine3.11 as builder

# Add Maintainer Info
LABEL maintainer="Suvarna Rokade"

# Environment Variables
ENV SERVER_URL=${SERVER_URL} 

# Install git
RUN apk --update add git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

FROM alpine  

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder go/src/app/ .

# Expose port 8090 to the outside world
EXPOSE 9090

RUN chmod a+x client

# Command to run the executable
CMD ["./client"]




