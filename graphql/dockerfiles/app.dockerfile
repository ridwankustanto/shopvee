FROM golang:1.17-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /github.com/ridwankustanto/shopvee
COPY go.mod go.sum ./
COPY account account
COPY product product
COPY graphql graphql
RUN GO111MODULE=on go build -o /go/bin/app ./graphql

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
