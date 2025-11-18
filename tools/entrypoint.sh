#!/bin/sh

URL=$POSTGRES_URL
goose -dir /migrations postgres "$URL" up

echo "Migrations completed"