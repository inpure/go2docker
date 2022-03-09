FROM golang:latest

RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go2docker"]