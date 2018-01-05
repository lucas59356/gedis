FROM golang:latest

ADD . /go/src/github.com/lucas59356/gedis
WORKDIR /go/src/github.com/lucas59356/gedis
RUN make core_test
RUN go build github.com/lucas59356/gedis/cli/gedis-api
EXPOSE 5935
CMD [ "./gedis-api" ]