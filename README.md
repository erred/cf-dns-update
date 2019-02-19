# cf-dns-update

[![Build](https://img.shields.io/badge/endpoint.svg?url=https://badger.seankhliao.com/r/github_seankhliao_cf-dns-update)](https://console.cloud.google.com/cloud-build/builds?project=com-seankhliao&query=source.repo_source.repo_name%20%3D%20%22github_seankhliao_cf-dns-update%22)
[![License](https://img.shields.io/github/license/seankhliao/cf-dns-update.svg?style=for-the-badge)](LICENSE)

Update a Cloudflare DNS record to the ip you're currently at

## Setup

Env vars to set:

- `ZONE_NAME`
  - set to zone name (domain)
- `RECORD_NAME`
  - set to the record name (sub(domain))
- `RECORD_TYPE`
  - set to `A` or `CNAME`
- `REPLICAS`
  - set to the number of simultaneous records to allow
- `CONTENT`
  - set to the domain for CNAMEs to point to
- `NO_PROXY`
  - set to `1` or `true` to bypass cloudflare's proxy
- `X_AUTH_EMAIL`
  - Email for cloudflare api
- `X_AUTH_KEY`
  - Key for cloudflare api

## External Dependencies

- Cloudflare
  - obviously
- ipify.org
  - get current ip
