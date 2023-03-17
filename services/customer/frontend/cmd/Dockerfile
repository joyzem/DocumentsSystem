FROM golang:alpine as builder

WORKDIR /app/services/customer/frontend/cmd

COPY /services/customer/frontend /app/services/customer/frontend
COPY /services/customer/dto /app/services/customer/dto
COPY /services/customer/domain /app/services/customer/domain
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

WORKDIR /app/services/customer/frontend/cmd

EXPOSE 8085

CMD ["./main"]
