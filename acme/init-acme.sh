#!/bin/bash
# Run from the acme/ directory: cd acme && ./init-acme.sh

set -e
set -u
set -o pipefail
set -x

# Wrapper: acme.sh exits 2 if cert is already valid and unchanged; treat that as success
ACME_RAW="docker-compose run --rm acme.sh acme.sh"
acme() { $ACME_RAW "$@" || [ $? -eq 2 ]; }
ACME=acme
NGINX="docker-compose -f ../docker-compose.yml exec nginx nginx"

# Use Let's Encrypt (acme.sh defaults to ZeroSSL as of v3)
$ACME --set-default-ca --server letsencrypt

# Create cert dirs (needed for fresh installs)
mkdir -p \
  ../data/certbot/conf/live/marktai.com \
  ../data/certbot/conf/live/wildcard.mtai.io \
  ../data/certbot/conf/live/wildcard.realisticcode.com \
  ../data/certbot/conf/live/24hlimes.com \
  ../data/certbot/conf/live/24hourlimes.com \
  ../data/certbot/conf/live/twentyfourhourlimes.com \

  ../data/certbot/conf/live/bonneyruan.com \
  ../data/certbot/conf/live/bonney.design \
  ../data/certbot/conf/live/millends.com


# --- Wildcard certs via Namecheap DNS challenge ---
# Note: *.marktai.com covers subdomains (www, clover, blog, etc.)
#       but NOT the naked domain, so we include marktai.com as a SAN.

echo "### Issuing *.marktai.com + marktai.com..."
$ACME --issue --dns dns_namecheap -d "*.marktai.com" -d "marktai.com"
$ACME --install-cert -d "*.marktai.com" \
  --fullchain-file /etc/letsencrypt/live/marktai.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/marktai.com/privkey.pem

echo "### Issuing *.mtai.io + mtai.io..."
$ACME --issue --dns dns_namecheap -d "*.mtai.io" -d "mtai.io"
$ACME --install-cert -d "*.mtai.io" \
  --fullchain-file /etc/letsencrypt/live/wildcard.mtai.io/fullchain.pem \
  --key-file /etc/letsencrypt/live/wildcard.mtai.io/privkey.pem

echo "### Issuing *.realisticcode.com + realisticcode.com..."
$ACME --issue --dns dns_namecheap -d "*.realisticcode.com" -d "realisticcode.com"
$ACME --install-cert -d "*.realisticcode.com" \
  --fullchain-file /etc/letsencrypt/live/wildcard.realisticcode.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/wildcard.realisticcode.com/privkey.pem


# --- Single-domain certs via webroot ---
# nginx must be running: docker-compose -f ../docker-compose.yml up -d nginx

echo "### Issuing 24hlimes.com..."
$ACME --issue --webroot /var/www/certbot -d "24hlimes.com"
$ACME --install-cert -d "24hlimes.com" \
  --fullchain-file /etc/letsencrypt/live/24hlimes.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/24hlimes.com/privkey.pem

echo "### Issuing 24hourlimes.com..."
$ACME --issue --webroot /var/www/certbot -d "24hourlimes.com"
$ACME --install-cert -d "24hourlimes.com" \
  --fullchain-file /etc/letsencrypt/live/24hourlimes.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/24hourlimes.com/privkey.pem

echo "### Issuing twentyfourhourlimes.com..."
$ACME --issue --webroot /var/www/certbot -d "twentyfourhourlimes.com"
$ACME --install-cert -d "twentyfourhourlimes.com" \
  --fullchain-file /etc/letsencrypt/live/twentyfourhourlimes.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/twentyfourhourlimes.com/privkey.pem


echo "### Issuing bonneyruan.com + www.bonneyruan.com..."
$ACME --issue --webroot /var/www/certbot -d "bonneyruan.com" -d "www.bonneyruan.com"
$ACME --install-cert -d "bonneyruan.com" \
  --fullchain-file /etc/letsencrypt/live/bonneyruan.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/bonneyruan.com/privkey.pem

echo "### Issuing bonney.design + www.bonney.design..."
$ACME --issue --webroot /var/www/certbot -d "bonney.design" -d "www.bonney.design"
$ACME --install-cert -d "bonney.design" \
  --fullchain-file /etc/letsencrypt/live/bonney.design/fullchain.pem \
  --key-file /etc/letsencrypt/live/bonney.design/privkey.pem

echo "### Issuing millends.com..."
$ACME --issue --webroot /var/www/certbot -d "millends.com"
$ACME --install-cert -d "millends.com" \
  --fullchain-file /etc/letsencrypt/live/millends.com/fullchain.pem \
  --key-file /etc/letsencrypt/live/millends.com/privkey.pem


echo "### Reloading nginx..."
$NGINX -s reload

echo "### Done. Start the daemon: docker-compose up -d"
