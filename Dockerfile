FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git make curl

WORKDIR /app

COPY . .

# Download all the dependencies
RUN go mod download
RUN go mod verify
RUN go mod tidy

# Install the package
RUN go install -v ./...

RUN export PATH=$PATH:/usr/local/go/bin
RUN go build -o bin/go-starter-template

#RUN go build -o bin-go-starter-template
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh
EXPOSE 9000

RUN chmod +x /app/entrypoint.sh
CMD ["sh", "-c", "/app/entrypoint.sh"]