# cf-dns-update

[![Build](https://img.shields.io/badge/endpoint.svg?url=https://badger.seankhliao.com/r/github_seankhliao_cf-dns-update)](https://console.cloud.google.com/cloud-build/builds?project=com-seankhliao&query=source.repo_source.repo_name%20%3D%20%22github_seankhliao_cf-dns-update%22)
[![License](https://img.shields.io/github/license/seankhliao/cf-dns-update.svg?style=for-the-badge)](LICENSE)

Update a Cloudflare DNS record to the ip you're currently at

## Setup

Env vars to set:

- `RECORD_NAME`: ex: sub.example.com
- `RECORD_ID`: Cloudflare DNS Record identifier
  - get using the list api
- `ZONE_ID`: Cloudflare Zone identifier
  - get from dash
- `X_AUTH_EMAIL`
  - Email for cloudflare api
- `X_AUTH_KEY`
  - Key for cloudflare api

## External Dependencies

- Cloudflare
  - obviously
- ipify.org
  - get current ip
