package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for AutoAR
type Config struct {
	// Telegram settings
	TelegramBotToken string
	TelegramChatID   string

	// Nuclei settings
	NucleiTemplatesPath string
	NucleiSeverity      string

	// Scanning settings
	Concurrency    int
	TimeoutSeconds int
	OutputDir      string

	// Notification settings
	NotifyOnNew      bool
	NotifyOnCritical bool
}

// Load reads configuration from environment variables (and optional .env file)
func Load(envFile string) (*Config, error) {
	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			return nil, fmt.Errorf("loading env file %q: %w", envFile, err)
		}
	}

	concurrency, err := strconv.Atoi(getEnv("CONCURRENCY", "5"))
	if err != nil {
		return nil, fmt.Errorf("invalid CONCURRENCY value: %w", err)
	}

	timeout, err := strconv.Atoi(getEnv("TIMEOUT_SECONDS", "30"))
	if err != nil {
		return nil, fmt.Errorf("invalid TIMEOUT_SECONDS value: %w", err)
	}

	cfg := &Config{
		TelegramBotToken:    os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChatID:      os.Getenv("TELEGRAM_CHAT_ID"),
		NucleiTemplatesPath: getEnv("NUCLEI_TEMPLATES_PATH", "/root/nuclei-templates"),
		NucleiSeverity:      getEnv("NUCLEI_SEVERITY", "critical,high,medium"),
		Concurrency:         concurrency,
		TimeoutSeconds:      timeout,
		OutputDir:           getEnv("OUTPUT_DIR", "/tmp/autoar-output"),
		NotifyOnNew:         getEnvBool("NOTIFY_ON_NEW", true),
		NotifyOnCritical:    getEnvBool("NOTIFY_ON_CRITICAL", true),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.TelegramBotToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN is required")
	}
	if c.TelegramChatID == "" {
		return fmt.Errorf("TELEGRAM_CHAT_ID is required")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return b
}
