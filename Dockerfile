FROM golang:latest
RUN go install ./... -v
CMD ["./projectOne"]
