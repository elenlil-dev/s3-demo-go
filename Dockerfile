FROM golang:1.23.6-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy

RUN go build  -o server ./cmd/server

FROM alpine:latest


WORKDIR /app/

COPY --from=builder /app/server .
COPY --from=builder /app/image ./image
COPY --from=builder /app/download ./download

EXPOSE 8080

CMD [ "./server" ]
