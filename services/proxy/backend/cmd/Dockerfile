FROM golang:alpine AS builder

WORKDIR /app/services/proxy/backend/cmd

COPY /services/proxy/backend /app/services/proxy/backend
COPY /services/proxy/domain /app/services/proxy/domain
COPY /services/proxy/dto /app/services/proxy/dto
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app/services/proxy/backend/cmd 

COPY --from=builder /app/services/proxy/backend/cmd/main .

EXPOSE 7076

CMD ["./main"]