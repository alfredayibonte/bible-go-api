FROM golang:1.22.1-alpine3.18
ENV DB_HOST db
ENV DB_USER test
ENV DB_PASS test
ENV DB_NAME test

WORKDIR /go/src/github.com/rkeplin/bible-go-api

COPY . .

RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["bible-go-api"]
