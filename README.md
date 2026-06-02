# DevStats Card

A small self-hosted service that renders your [CNCF DevStats](https://devstats.cncf.io) contributions as an SVG, ready to embed in a GitHub profile README.

[![DevStats](https://devstats.app/?username=skarlso)](https://github.com/Skarlso/devstats)

## Usage

Add this to your README and swap in your GitHub username:

```markdown
[![DevStats](https://devstats.app/?username=skarlso)](https://github.com/Skarlso/devstats)
```

The server queries the CNCF DevStats API and returns an `image/svg+xml` card showing your DevStats score, pull requests, and issues. Responses are cached for two hours.

### Themes

Add `&theme=` to pick a palette. Defaults to `dark`; an unknown value falls back to `dark`.

```markdown
[![DevStats](https://devstats.app/?username=skarlso&theme=light)](https://github.com/Skarlso/devstats)
```

| Theme   | Description            |
|---------|------------------------|
| `dark`  | Dark gradient (default). |
| `light` | Light, GitHub-light palette. |
| `cncf`  | Purple gradient. |

## Endpoints

| Method | Path                             | Description              |
|--------|----------------------------------|--------------------------|
| GET    | `/?username=<id>&theme=<theme>`  | Render the card.         |
| GET    | `/health`                        | Liveness check, `OK`.    |

## Running it yourself

The SVG template is embedded in the binary, so the build is fully self-contained. No assets to ship alongside it.

```bash
go build -trimpath -ldflags="-s -w" -o devstats .
./devstats
```

The server listens on `127.0.0.1:8080`. Cross-compile for a Raspberry Pi (arm64) with:

```bash
GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w" -o devstats .
```

By default it binds to `127.0.0.1:8080`. Override with `LISTEN_ADDR` (for example `:8080` inside a container, then publish the port only to the host loopback). Run it behind a reverse proxy or a [Cloudflare Tunnel](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/) so there are no open ports and TLS terminates at the edge.

## Credits

Based on the original [tico88612/devstats-card](https://github.com/tico88612/devstats-card) which was abandoned.
