#!/bin/bash
# wait-for-it.sh
# Wait for a service to be available before continuing

set -e

host="$1"
port="$2"
timeout="${3:-15}" # Default timeout of 15 seconds

echo "Waiting for $host:$port to be available..."
echo $host:$port
while ! nc -z "$host" "$port" >/dev/null 2>&1; do
  timeout=$((timeout - 1))
  if [ "$timeout" -le 0 ]; then
    echo "Error: Timeout. $host:$port is not available."
    exit 1
  fi
  sleep 1
done

echo "$host:$port is available. Continuing with the execution."
