########################################################
# STEP 1 build executable binary
########################################################
FROM golang:alpine AS builder

# Install git for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/notes/app/

COPY . .

# Fetch dependencies using go get.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/notes-api-server ./cmd/main.go

########################################################
# STEP 2 build image
########################################################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/notes-api-server /go/bin/notes-api-server

EXPOSE 8080/tcp

# Run the hello binary.
ENTRYPOINT ["/go/bin/notes-api-server"]