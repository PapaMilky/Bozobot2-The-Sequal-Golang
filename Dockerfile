FROM golang:1.18-alpine

ENV TOKEN=""
ENV GUILD=""

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o botexecutable

CMD GUILD_ID="${GUILD}" BOT_TOKEN="${TOKEN}" go run .