#!/bin/bash

case "$1" in
  up)
    echo "🚀 Starting docker-compose..."
    docker-compose up -d
    ;;
  down)
    echo "🛑 Stopping docker-compose..."
    docker-compose down
    ;;
  restart)
    echo "🔄 Restarting docker-compose..."
    docker-compose down
    docker-compose up -d
    ;;
  *)
    echo "Usage: $0 {up|down|restart}"
    exit 1
    ;;
esac
