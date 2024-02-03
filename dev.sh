#!/bin/sh

cleanup() {
    kill $(lsof -t -i:3000)
    pkill -f "air -c ./.air.toml"
    exit 0
}

# Trap the INT signal (Ctrl+C) and call cleanup
trap cleanup INT

air -c ./.air.toml \
    & npx tailwindcss -i ./styles.css -o ./public/css/app.css --watch \
    & browser-sync start \
    --files './templates/**/*.html, ./public/**/*.css, ./public/**/*.js, ./markdown/**/*.md, ./tmp/main' \
    --port 3001 \
    --proxy 'localhost:3000' \
    --no-open \
    --middleware 'function(req, res, next) { \
        res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
        return next(); \
    }'