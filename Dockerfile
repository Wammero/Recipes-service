FROM golang:1.21.1

COPY . /

WORKDIR /

RUN go build main.go

CMD ["./main"]

EXPOSE 8080