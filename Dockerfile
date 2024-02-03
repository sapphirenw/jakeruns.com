# Node stage to handle Tailwind and JavaScript minification
FROM node:current-alpine as node-builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install Node dependencies including UglifyJS
RUN npm install && npm install uglify-js -g

# Copy Tailwind and JavaScript source files
COPY styles.css .
COPY tailwind.config.js .
COPY public/js/ ./public/js/

# Generate minified CSS file using Tailwind
RUN npx tailwindcss -o ./public/css/app.css

# Minify JavaScript files
RUN find ./public/js -name '*.js' -exec uglifyjs --compress --mangle --output {} -- {} \;

# Go stage to build the Go application
FROM golang:alpine as go-builder

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the Go source files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage to prepare the runtime image
FROM scratch

# Copy the CA certificates
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the Go binary
COPY --from=go-builder /app/main .

# Copy minified CSS from node-builder
COPY --from=node-builder /app/public/css/app.css /public/css/

# Copy static assets and templates
COPY --from=go-builder /app/public /public
COPY --from=go-builder /app/templates /templates

# Expose the port the server listens on
EXPOSE 3000

# Run the compiled binary
CMD ["/main"]
