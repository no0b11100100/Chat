FROM golang:latest

WORKDIR /Chat/Server

ADD . .

RUN go mod download
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17

RUN go build -o /server
EXPOSE 8080

CMD ["/server"]
