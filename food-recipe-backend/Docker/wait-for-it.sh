#!/usr/bin/env bash
# Wait-for-it script to check service availability

TIMEOUT=30
while ! nc -z db 5432; do   
  echo "Waiting for PostgreSQL..."
  sleep 1
  TIMEOUT=$((TIMEOUT-1))
  if [ "$TIMEOUT" -le 0 ]; then
    echo "PostgreSQL not available after 30s, exiting."
    exit 1
  fi
done

echo "PostgreSQL is available!"
exec "$@"
