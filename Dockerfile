FROM golang:1.116-alpine3.13 as builder
COPY go.mod go.sum /go/src/github.com/harykr/ticket-booking-system/
WORKDIR /go/src/github.com/harykr/ticket-booking-system
RUN go mod download
COPY . /go/src/github.com/harykr/ticket-booking-system
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bookingsystem github.com/harykr/ticket-booking-system

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/harykr/ticket-booking-system/build/bookingsystem /usr/bin/bookingsystem
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/bookingsystem"]