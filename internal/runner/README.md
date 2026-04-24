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
|---------------|-------------------------|------------------------------------|
| `Concurrency` | `config.Config` / `CONCURRENCY` env | Number of parallel workers |

## Extending

Replace the `execute` function body in `runner.go` with real tool invocations
(e.g. `subfinder`, `nmap`) to integrate external recon tools.
