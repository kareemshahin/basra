#!/usr/bin/env bash

URL="http://localhost:8000/games"

echo "testing POST games"
curl -H "Content-Type: application/json" \
  -X POST \
  -d '{"name":"my test game"}' \
  "${URL}"

echo "testing GET games"
curl -i \
  -H "Accept: application/json" \
  "${URL}"

