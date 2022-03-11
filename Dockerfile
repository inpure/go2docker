FROM golang:latest

COPY ./main.go /opt
RUN cd /opt && go mod tidy && go build ./main.go
EXPOSE 54321
ENTRYPOINT ["/opt/main"]