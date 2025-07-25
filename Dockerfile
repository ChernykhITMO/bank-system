FROM golang:1.23

WORKDIR /bankSystem

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]