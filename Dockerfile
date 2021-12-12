FROM golang:alpine

WORKDIR /go/src/argos
COPY . .

RUN go mod download

CMD ["go", "run", "./main.go"]
