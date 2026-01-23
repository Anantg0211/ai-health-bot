#!/bin/sh
set -x

# Define environment variables
API_ENV="development"
API_KEY=""
SERVICE_NAME="ai-powered-health-bot"
CERT_SERVICE="certs-service"
CERT_API_KEY=${CERT_API_KEY:-""}
# Ensure the config directory exists
mkdir -p /config

mkdir -p /downloads
chmod 777 /downloads

# Execute the command passed to the script
exec "$@"