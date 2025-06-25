FROM debian:bookworm

# install nftables and netcat for debugging
RUN apt-get update && apt-get install -y nftables netcat-openbsd

CMD ["tail", "-f", "/dev/null"]