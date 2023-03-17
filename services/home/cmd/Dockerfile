FROM golang:alpine as builder

WORKDIR /app/services/home/cmd

COPY /services/home /app/services/home
COPY /services/base /app/services/base
COPY go.mod /app
COPY go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

WORKDIR /app/services/home/cmd

EXPOSE 80

CMD ["./main"]
