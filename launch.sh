#!/bin/bash

# Set environment variables
export MINIFLUX_URL="https://your-miniflux-instance.com"
export MINIFLUX_USER="your-miniflux-username"
export MINIFLUX_PASS="your-miniflux-password"
export GMAIL_EMAIL="your-gmail-email@example.com"
export GMAIL_PASSWORD="your-gmail-password"
export CATEGORY="daily"

# Build the application
go build

# Run the application
./miniflux-email-updates
