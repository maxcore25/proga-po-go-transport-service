#!/bin/sh
set -e

echo "==> Starting Go REST API on :9000 ..."
/app/server &
GO_PID=$!

# Ждём, пока API поднимется (максимум 15 секунд)
echo "==> Waiting for Go API to be ready..."
for i in $(seq 1 15); do
    if wget -q --spider http://127.0.0.1:9000/health 2>/dev/null; then
        echo "==> Go API is up."
        break
    fi
    sleep 1
done

echo "==> Starting Nginx ..."
nginx -g "daemon off;" &
NGINX_PID=$!

# Graceful shutdown: при получении SIGTERM/SIGINT останавливаем оба процесса
trap 'echo "==> Shutting down..."; kill $NGINX_PID $GO_PID; wait' TERM INT

# Мониторим оба процесса — если один упал, убиваем второй и выходим
wait $GO_PID
echo "==> Go API exited unexpectedly. Stopping Nginx..."
kill $NGINX_PID
wait $NGINX_PID
