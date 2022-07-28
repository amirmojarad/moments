FROM golang:alpine

LABEL maintainer="moments"

RUN apk add --no-cache git && apk add --no-cache bash && apk add build-base

Run mkdir /app
WORKDIR /app

COPY . .
COPY /env .
COPY .env .
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]