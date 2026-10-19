# nanoflux-plugin-patreon

A nanoflux external plugin for Patreon creator feeds. It reads public posts
through Patreon's web API.

## Install

```bash
cd nanoflux/plugins/patreon
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../nanoflux-plugin-patreon .
```

Restart nanoflux after putting the binary in its plugins directory
(`NF_PLUGINS_DIR`, `/plugins` in the container).

## Usage

Add a creator page, for example `https://www.patreon.com/<creator>`.
