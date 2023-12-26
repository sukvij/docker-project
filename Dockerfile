FROM golang:1.20.5-alpine3.18
WORKDIR /app
COPY . /app
RUN go build /app
EXPOSE 8000
ENTRYPOINT [ "./vijju" ]