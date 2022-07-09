FROM golang:1.18.3-alpine AS builder

RUN apk update && apk add --no-cache --update git curl bash

RUN mkdir /app
WORKDIR /app

#RUN export GO111MODULE=on

# install reloader
RUN go install github.com/cespare/reflex@latest

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

# RUN go build -o ./run .
RUN go build

FROM alpine
WORKDIR /app
# COPY --from=builder /app/run .
COPY --from=builder /app/golang .

EXPOSE 8080
# CMD ["./run"]
CMD ["./golang"]
