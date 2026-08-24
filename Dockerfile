#stage 1
FROM node:20-alpine AS tailwind-builder

WORKDIR /app

RUN npm install -g tailwindcss @tailwindcss/cli

ENV NODE_PATH=/usr/local/lib/node_modules

COPY . .

RUN tailwindcss -i ./src/input.css -o ./src/output.css --minify

# stage 2
FROM golang:alpine3.22 AS builder

WORKDIR /crescendo

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/api

# Stage 3

FROM alpine:3.22

WORKDIR /root

COPY --from=builder /app /root/

CMD [ "./app" ]

#CMD [ "air" ]