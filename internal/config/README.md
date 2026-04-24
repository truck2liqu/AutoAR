# Config Package

The `config` package is responsible for loading and validating all runtime
configuration for **AutoAR** from environment variables (or an optional `.env`
file).

## Usage

```go
import "github.com/your-org/autoar/internal/config"

cfg, err := config.Load(".env") // pass "" to skip file loading
if err != nil {
    log.Fatalf("config error: %v", err)
}
```

## Environment Variables

| Variable | Required | Default | Description |
|---|---|---|---|
| `TELEGRAM_BOT_TOKEN` | ✅ | — | Telegram bot token for notifications |
| `TELEGRAM_CHAT_ID` | ✅ | — | Target Telegram chat / channel ID |
| `NUCLEI_TEMPLATES_PATH` | ❌ | `/root/nuclei-templates` | Path to nuclei templates directory |
| `NUCLEI_SEVERITY` | ❌ | `critical,high,medium` | Comma-separated severity filter |
| `CONCURRENCY` | ❌ | `5` | Number of concurrent scan workers |
| `TIMEOUT_SECONDS` | ❌ | `30` | Per-request timeout in seconds |
| `OUTPUT_DIR` | ❌ | `/tmp/autoar-output` | Directory for scan result files |
| `NOTIFY_ON_NEW` | ❌ | `true` | Send notification for every new finding |
| `NOTIFY_ON_CRITICAL` | ❌ | `true` | Always notify for critical severity |

## Validation

`Load` returns an error if any **required** variable is missing or if a
numeric variable cannot be parsed. The caller should treat this as a fatal
startup error.
