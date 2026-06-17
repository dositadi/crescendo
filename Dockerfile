FROM golang:alpine3.22 AS builder

WORKDIR /crescendo

#RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/api

FROM alpine:3.22

WORKDIR /root

COPY --from=builder /app /root/

CMD [ "./app" ]

#CMD [ "air" ]