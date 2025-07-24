FROM golang:1.24

WORKDIR /app

# Install dependencies untuk sqlite3
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

ENV CGO_ENABLED=1

# Copy dependency info
COPY go.mod ./

RUN ls -al
RUN go mod download

# Copy seluruh aplikasi
COPY . .

# Build binary
RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
