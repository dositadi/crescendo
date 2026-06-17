FROM golang:alpine3.22 AS builder

WORKDIR /groupie-tracker

#RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/

FROM alpine3.22

WORKDIR /root

COPY --from=builder /app /root/

CMD [ "./app" ]

#CMD [ "air" ]