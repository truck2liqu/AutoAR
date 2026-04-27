# Runner

The `runner` package provides a concurrent task execution engine for AutoAR.

## Overview

A `Runner` manages a pool of worker goroutines that consume `Task` values from
an internal channel and emit `Result` values once each task completes.

## Usage

```go
cfg, _ := config.Load()
r := runner.New(cfg)
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

r.Start(ctx)

// Submit tasks
r.Submit(runner.Task{Target: "example.com", Type: "subdomain"})

// Stop accepting new tasks and wait for workers
r.Stop()

// Consume results
for res := range r.Results() {
    if res.Err != nil {
        log.Printf("error: %v", res.Err)
        continue
    }
    fmt.Println(res.Output)
}
```

## Configuration

| Field         | Source                  | Description                        |
|---------------|-------------------------|-------------------------------------|
| `Concurrency` | `config.Config` / `CONCURRENCY` env | Number of parallel workers (default: 5) |

> **Personal note:** I've found that setting `CONCURRENCY=10` works well on my
> machine for most recon workloads without hammering rate limits. Adjust down
> to 3–5 if you start seeing timeouts from external tools.
>
> On my home lab (8-core machine, ~100Mbps uplink) I've settled on `CONCURRENCY=8`
> as a sweet spot — fast enough to feel snappy, conservative enough that nuclei
> doesn't get rate-limited by most bug bounty targets.

## Extending

Replace the `execute` function body in `runner.go` with real tool invocations
(e.g. `subfinder`, `nmap`) to integrate external recon tools.

### Tools I plan to integrate

- `subfinder` – passive subdomain enumeration
- `httpx` – HTTP probing
- `nuclei` – vulnerability scanning (run last, after scope is confirmed)
- `katana` – crawling / JS endpoint discovery (added to my list after finding it useful on a recent program)
