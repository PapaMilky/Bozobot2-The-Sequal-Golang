FROM golang:1.18-alpine

ENV TOKEN=""
ENV GUILD=""

WORKDIR /app

RUN apk update
RUN apk add git

COPY go.mod ./
COPY go.sum ./
COPY ./utils /go/src/Bozobot2-The-Sequal-Golang/
RUN ls
RUN go mod download


COPY . ./

RUN go build -buildvcs=false -o botexecutable

CMD GUILD_ID="${GUILD}" BOT_TOKEN="${TOKEN}" go run . -buildvcs=false