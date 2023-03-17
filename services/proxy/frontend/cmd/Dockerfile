FROM golang:alpine as builder

WORKDIR /app/services/proxy/frontend/cmd

COPY /services/proxy/frontend /app/services/proxy/frontend
COPY /services/proxy/dto /app/services/proxy/dto
COPY /services/proxy/domain /app/services/proxy/domain
COPY /services/employee/dto /app/services/employee/dto
COPY /services/employee/domain /app/services/employee/domain
COPY /services/organization/dto /app/services/organization/dto
COPY /services/organization/domain /app/services/organization/domain
COPY /services/product/dto /app/services/product/dto
COPY /services/product/domain /app/services/product/domain
COPY /services/account/dto /app/services/account/dto
COPY /services/account/domain /app/services/account/domain
COPY /services/documents /app/services/documents
COPY /services/customer/dto /app/services/customer/dto
COPY /services/customer/domain /app/services/customer/domain
COPY /services/base /app/services/base
COPY /go.mod /app
COPY /go.sum /app

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

WORKDIR /app/services/proxy/frontend/cmd

EXPOSE 8086

CMD ["./main"]
