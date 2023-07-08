#!/bin/bash

set -e

# Set the migration parameters
databaseUrl="postgres://meaning:password@localhost:5432/meaning?sslmode=disable"
migrationPath="./migrations"

if [ "$1" == "down" ]; then
  # Run the migration down
  migrate -database "$databaseUrl" -path "$migrationPath" down
else
  # Run the migration up
  migrate -database "$databaseUrl" -path "$migrationPath" up
fi
