#!/bin/sh
# wait-for.sh

set -e

host="$1"
shift
cmd="$@"

until pg_isready -h "$host" -p 5432 -U "$DB_USER"; do
  echo "⏳ Waiting for PostgreSQL at $host:5432..."
  sleep 1
done

echo "✅ PostgreSQL is ready, running: $cmd"
exec $cmd
