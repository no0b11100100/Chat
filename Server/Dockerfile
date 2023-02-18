FROM golang:latest

WORKDIR /Chat/Server

ADD . .

RUN go mod download
RUN go mod tidy -go=1.18

RUN go build -o /server
EXPOSE 8080

CMD ["/server"]
