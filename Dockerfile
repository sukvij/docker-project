FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go get github.com/gin-gonic/gin

RUN go build -o /go-docker-demo

EXPOSE 8000

CMD [ "/go-docker-demo" ]