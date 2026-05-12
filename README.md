# crawler

Minimal Go web crawler built in Go.

It visits pages within the same domain, extracts useful page data, and writes the results to `report.json`.

## What It Collects

- Page URL
- Heading
- First paragraph
- Outgoing links
- Image URLs

## How It Works

- Starts from a base URL
- Crawls pages concurrently
- Stays on the same host
- Stops after the max page limit is reached
- Normalizes URLs to avoid duplicate visits

## Usage

### Install

```bash
git clone <repo-url>
cd crawler
go mod download
```

```bash
go run . <baseURL> <maxConcurrency> <maxPages>
```

Example:

```bash
go run . https://example.com 5 50
```

## Output

The crawler writes a sorted JSON report to `report.json` in the project root.
