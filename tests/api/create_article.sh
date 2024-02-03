#!/bin/sh

curl -X POST \
    -H "x-api-key: $API_KEY" \
    -H "Content-Type: application/json" \
    --data @./create_article.json \
    localhost:3000/api/articles