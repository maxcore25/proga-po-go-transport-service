# ─────────────────────────────────────────────
# Stage 1: Generate TLS certificates
# ─────────────────────────────────────────────
FROM alpine:3.19 AS certs

RUN apk add --no-cache openssl

WORKDIR /certs

# Self-signed cert for localhost (SAN включён)
RUN openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
    -keyout server.key \
    -out server.crt \
    -subj "/CN=localhost/O=MyApp/C=KZ" \
    -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"


# ─────────────────────────────────────────────
# Stage 2: Build Go binary
# ─────────────────────────────────────────────
FROM golang:1.25.4-alpine AS builder

# Зависимости для CGO (если нужен sqlite и т.п.)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Кэшируем зависимости отдельным слоем
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходник
COPY . .

# Генерируем Swagger-документацию
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g cmd/app/main.go -o ./docs --parseInternal --parseDependency || true && \
    go get github.com/maxcore25/proga-po-go-transport-service/docs && \
    go mod tidy

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/server ./cmd/app

# ─────────────────────────────────────────────
# Stage 3: Final image — Go app + Nginx
# ─────────────────────────────────────────────
FROM nginx:1.25-alpine

# ── Системные пакеты ──────────────────────────
RUN apk add --no-cache ca-certificates tzdata

# ── Go-бинарник ───────────────────────────────
COPY --from=builder /app/server /app/server

# Если есть статические файлы / шаблоны / миграции — копируем их тоже
COPY --from=builder /app/internal/migrations /app/internal/migrations

# ── Swagger docs (генерируются swag init) ─────
COPY --from=builder /app/docs /app/docs

# ── TLS-сертификаты ───────────────────────────
COPY --from=certs /certs/server.crt /etc/nginx/ssl/server.crt
COPY --from=certs /certs/server.key /etc/nginx/ssl/server.key

# ── Конфиг Nginx ──────────────────────────────
COPY docker/nginx/nginx.conf /etc/nginx/nginx.conf

# ── Стартовый скрипт ──────────────────────────
COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Рабочая директория для приложения
WORKDIR /app

# Открываем только HTTPS-порт
EXPOSE 8888

ENTRYPOINT ["/entrypoint.sh"]
