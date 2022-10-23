FROM golang:latest

WORKDIR /Chat/Client/Server

ADD . .

RUN go mod download

# RUN go get github.com/mailru/easyjson@latest

RUN go mod tidy
RUN go build -o /client-server
EXPOSE 8080

CMD ["/client-server"]
