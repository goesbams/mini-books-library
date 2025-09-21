FROM golang:1.25.1-alpine as builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod tidy

COPY ./backend/ /app/backend/

COPY ./backend/config/config.dev.yaml /app/backend/config/config.dev.yaml

WORKDIR /app/backend/cmd

RUN go build -o backend .

FROM alpine:latest

RUN apk add --no-cache wget \
    && wget -q https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz \
    && tar -xvzf migrate.linux-amd64.tar.gz \
    && mv migrate /usr/local/bin/

COPY --from=builder /app/backend/cmd/backend /usr/local/bin/backend

COPY --from=builder /app/backend/config/config.dev.yaml /app/backend/config/config.dev.yaml
COPY ./backend/migrations /app/backend/migrations

COPY ./backend/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 9000

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
