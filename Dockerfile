FROM golang:1.24.2-alpine AS builder

WORKDIR /build

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 9000

CMD [ "/build/main" ]