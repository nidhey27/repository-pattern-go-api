#!/bin/bash

# Loop 100 times
for i in {1..100}; do
  # Your curl command
  curl --location 'http://127.0.0.1:8080/api/user' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30,
    "birthday": "1992-01-15T00:00:00Z",
    "member_number": "123456",
    "activated_at": "2023-01-01T12:30:45Z"
  }'
done
