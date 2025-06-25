FROM golang:1.24

# install nftables and netcat for debugging
RUN apt-get update && apt-get install -y nftables netcat-openbsd

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

# run the air tool to watch for changes
CMD ["air", "--build.cmd", "go build -buildvcs=false -o /usr/local/bin/efw ./cmd/efw/", "--build.bin", "/usr/local/bin/efw", "--build.args_bin", "status"]