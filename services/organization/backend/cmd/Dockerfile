FROM golang:alpine AS builder

WORKDIR /app/services/organization/backend/cmd

COPY /services/organization/backend /app/services/organization/backend
COPY /services/organization/domain /app/services/organization/domain
COPY /services/organization/dto /app/services/organization/dto
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app/services/organization/backend/cmd 

COPY --from=builder /app/services/organization/backend/cmd/main .

EXPOSE 7072

CMD ["./main"]