#!/bin/bash

# Set environment variables
export MINIFLUX_URL="https://reader.miniflux.app/v1"
export MINIFLUX_TOKEN="token-miniflux"
export SMTP_USERNAME="your-gmail-email"

export SMTP_PASSWORD="your-password"
export SMTP_HOST="smtp.gmail.com"
export CATEGORY="daily"
export RECEIVER_EMAIL="contact@skatkov.com"

# Build the application
go build

# Run the application
go run .
