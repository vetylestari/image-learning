# Start from golang base image
FROM golang:1.18-alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache alpine-sdk

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN export PATH=$PATH:/usr/local/go/bin

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

#Setup hot-reload for dev stage
RUN go install github.com/githubnemo/CompileDaemon@latest

#RUN CompileDaemon -build="go build -buildvcs=false -o go-starter-template" -command="./go-starter-template"