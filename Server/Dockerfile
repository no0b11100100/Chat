FROM rust:latest

WORKDIR /Chat/Server

ADD . .
RUN cargo build --release
EXPOSE 8080

CMD ["./target/release/Server"]
