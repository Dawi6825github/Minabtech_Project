#!/bin/sh

# Wait for PostgreSQL to be ready
echo "Waiting for database..."
/wait-for-it.sh db:5432 --timeout=60 --strict -- echo "Database is up!"

# Run the backend application
echo "Starting backend service..."
exec ./backend
