FROM golang:1.25.1-alpine as builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod tidy

COPY ./backend/ /app/backend/

COPY ./backend/config/config.dev.yaml /app/backend/config/config.dev.yaml

WORKDIR /app/backend/cmd

RUN go build -o backend .

FROM alpine:latest

COPY --from=builder /app/backend/cmd/backend /usr/local/bin/backend

COPY --from=builder /app/backend/config/config.dev.yaml /app/backend/config/config.dev.yaml

EXPOSE 9000

CMD ["backend"]
