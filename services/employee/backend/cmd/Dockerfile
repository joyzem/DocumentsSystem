FROM golang:alpine as builder

WORKDIR /app/services/employee/backend/cmd

COPY /services/employee/backend /app/services/employee/backend
COPY /services/employee/dto /app/services/employee/dto
COPY /services/employee/domain /app/services/employee/domain
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app/services/employee/backend/cmd

COPY --from=builder /app/services/employee/backend/cmd/main .

EXPOSE 7074

CMD ["./main"]
