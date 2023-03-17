FROM golang:alpine as builder

WORKDIR /app/services/account/backend/cmd

COPY /services/account/backend /app/services/account/backend
COPY /services/account/dto /app/services/account/dto
COPY /services/account/domain /app/services/account/domain
COPY /services/base /app/services/base
COPY go.mod /app
COPY go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app/services/account/backend/cmd

COPY --from=builder /app/services/account/backend/cmd/main .

EXPOSE 7073

CMD ["./main"]
