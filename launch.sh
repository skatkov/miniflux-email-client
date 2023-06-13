#!/bin/bash

# Set environment variables
export MINIFLUX_URL="https://reader.miniflux.app/v1"
export MINIFLUX_TOKEN="token-miniflux"
export GMAIL_EMAIL="your-gmail-email"

export GMAIL_PASSWORD="your-password"
export CATEGORY="daily"
export RECEIVER_EMAIL="contact@skatkov.com"

# Build the application
go build

# Run the application
go run .
