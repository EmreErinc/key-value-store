FROM golang:alpine3.13

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

RUN mkdir -p /tmp
RUN mkdir -p /tmp/store

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# generate documentations
RUN apk update && \
    apk upgrade && \
    apk add --update alpine-sdk && \
    apk add --no-cache bash make
RUN make docs

# run tests
RUN make tests

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/*.env .
RUN cp /build/main .

# Export necessary port
EXPOSE 9090

# Command to run when starting the container
CMD ["/dist/main"]