# Use the official Golang image as the builder
FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

# RUN go get github.com/cospare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

# Use a minimal Alpine image for the final container
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the executable from the builder
COPY --from=builder /go/src/app/run .

EXPOSE 8000
CMD ["./run"]