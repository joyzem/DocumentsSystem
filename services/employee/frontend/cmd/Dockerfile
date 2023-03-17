FROM golang:alpine as builder

WORKDIR /app/services/employee/frontend/cmd

COPY /services/employee/frontend /app/services/employee/frontend
COPY /services/employee/dto /app/services/employee/dto
COPY /services/employee/domain /app/services/employee/domain
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

WORKDIR /app/services/employee/frontend/cmd

EXPOSE 8084

CMD ["./main"]
