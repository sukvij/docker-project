FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/go-docker-demo

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/go-docker-demo .


EXPOSE 8000

CMD [ "/go-docker-demo" ]