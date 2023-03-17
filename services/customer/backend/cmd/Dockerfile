FROM golang:alpine as builder

WORKDIR /app/services/customer/backend/cmd

COPY /services/customer/backend /app/services/customer/backend
COPY /services/customer/dto /app/services/customer/dto
COPY /services/customer/domain /app/services/customer/domain
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app/services/customer/backend/cmd

COPY --from=builder /app/services/customer/backend/cmd/main .

EXPOSE 7075

CMD ["./main"]
