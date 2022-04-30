FROM golang:latest

WORKDIR /Chat/Client/Server

ADD . .
RUN go mod download
RUN chmod +x install_dependencies.sh
RUN ./install_dependencies.sh
RUN go mod tidy

RUN go build -o /cliet-server
EXPOSE 8080

CMD ["/client-server"]
