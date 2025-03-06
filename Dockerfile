FROM golang:1.24

RUN apt-get update && apt-get install -y sqlite3 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p data
RUN sqlite3 data/bot.db "PRAGMA user_version;"

CMD ["air"]
