FROM golang:alpine as builder

WORKDIR /app/services/organization/frontend/cmd

COPY /services/organization/frontend /app/services/organization/frontend
COPY /services/organization/dto /app/services/organization/dto
COPY /services/organization/domain /app/services/organization/domain
COPY /services/account/dto /app/services/account/dto
COPY /services/account/domain /app/services/account/domain
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

WORKDIR /app/services/organization/frontend/cmd

EXPOSE 8082

CMD ["./main"]
