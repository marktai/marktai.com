#!/bin/bash
# Run from the acme/ directory: cd acme && ./rotate-ssl.sh

set -e
set -u
set -o pipefail

echo "### Force-renewing all certs..."
docker-compose run --rm acme.sh acme.sh --renew-all --force

echo "### Reloading nginx..."
docker-compose -f ../docker-compose.yml exec nginx nginx -s reload

echo "### Done."
