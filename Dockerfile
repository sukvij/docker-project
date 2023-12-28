FROM golang:latest

WORKDIR /app
COPY . ./
COPY ./go.mod ./
RUN go build -o main
CMD [ "./main" ]
