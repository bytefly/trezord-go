FROM golang:1.11.2

RUN mkdir /trezord-go
WORKDIR /trezord-go
COPY ./scripts/run_in_docker.sh /trezord-go

RUN apt-get update
RUN apt-get install -y redir

RUN go get github.com/bytefly/trezord-go
RUN go build github.com/bytefly/trezord-go

ENTRYPOINT '/trezord-go/run_in_docker.sh'
EXPOSE 11325