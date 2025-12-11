#!/bin/sh
set -e

# Wait for database to be ready
until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do
  echo "Waiting for database..."
  sleep 2
done

# Run migrations
echo "Running migrations..."
sql-migrate up -config=dbconfig.yml -env=development

# Start the server
echo "Starting server..."
exec /app/server
