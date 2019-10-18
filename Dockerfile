FROM golang:latest
RUN go install ./... -v
EXPOSE 8000
CMD ["./projectOne"]
