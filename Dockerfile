FROM golang:latest
ADD . /go/src/github.com/msarajam/projectOne
WORKDIR /go/src/github.com/msarajam/projectOne
RUN go install -v
CMD ["./projectOne"]