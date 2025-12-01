#!/bin/bash

set -e

# Config
SCHEMA_FILE="schema.sql"
DEV_URL="${DEV_URL:-postgres://sqlc:sqlcpass@localhost:5432/sqlcdb?sslmode=disable}"
FORMAT="{{ sql . \"  \" }}"

function migrate() {
 MIGRATION_NAME=${1:-"initial"}
  echo "🔧 Generating migration diff..."
  atlas migrate diff "$MIGRATION_NAME" \
    --to "file://$SCHEMA_FILE" \
    --dev-url "$DEV_URL" \
    --format "$FORMAT"
  echo "✅ Migration diff saved to $SCHEMA_FILE"
}

function apply() {
  echo "🔧 Applying migrations..."
  atlas migrate apply --url "$DEV_URL"
  echo "✅ Migrations applied successfully"
}

function status() {
  atlas migrate status --url "$DEV_URL"
}

function help() {
  echo "Usage: $0 {migrate|apply|status}"
}

# Main
case "$1" in
  migrate)
    migrate
    ;;
  apply)
    apply
    ;;
  status)
    status
    ;;
  *)
    help
    exit 1
    ;;
esac
