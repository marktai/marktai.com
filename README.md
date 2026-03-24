# marktai.com
Source code for my personal website

This repository contains both code for my Go server and the Angular Bootstrap front end.


## SSL certificates

Certs are managed by [acme.sh](https://github.com/acmesh-official/acme.sh) running as a Docker container. It auto-renews all certs in the background.

### Setup (first time or new server)

```
cd acme
cp .env.example .env        # fill in Namecheap API credentials
docker-compose -f ../docker-compose.yml up -d nginx
./init-acme.sh
docker-compose up -d        # start acme.sh daemon
```

### Force-renew all certs

```
cd acme && ./rotate-ssl.sh
```

