# Multi stage build
FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

WORKDIR /cab-service

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .


FROM scratch

COPY --from=builder /cab-service/bin/main .
EXPOSE 8000

# Run executable
CMD ["./main"]