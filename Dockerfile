FROM golang:1.19

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

COPY db.txt ./

RUN go build -o /docker-scores

CMD ["/docker-scores"]