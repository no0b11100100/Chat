FROM golang:latest

WORKDIR /Chat/RemoteServer

ADD . .
RUN go mod download
RUN chmod +x install_dependencies.sh
RUN ./install_dependencies.sh
RUN go mod tidy

RUN go build -o /server
EXPOSE 8080

CMD ["/server"]