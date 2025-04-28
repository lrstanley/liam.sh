// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

type Flags struct {
	HTTP     ConfigHTTP     `group:"HTTP Server options" namespace:"http"     env-namespace:"HTTP"`
	Database ConfigDatabase `group:"Database options"    namespace:"database" env-namespace:"DATABASE"`
	Github   ConfigGithub   `group:"Github options"      namespace:"github"   env-namespace:"GITHUB"`
	WakAPI   ConfigWakAPI   `group:"WakAPI options"      namespace:"wakapi"   env-namespace:"WAKAPI"`
	AI       ConfigAI       `group:"AI options"         namespace:"ai"       env-namespace:"AI"`
}

// ConfigDatabase holds the database configuration.
type ConfigDatabase struct {
	URL string `env:"URL" long:"url" default:"file:local.db" description:"database connection url"`
}

// ConfigHTTP are configurations specifically utilized by the HTTP service.
type ConfigHTTP struct {
	BaseURL        string   `env:"BASE_URL"        long:"base-url"        default:"http://localhost:8080" description:"base url for the HTTP server"`
	BindAddr       string   `env:"BIND_ADDR"       long:"bind-addr"       default:":8080" required:"true" description:"ip:port pair to bind to"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" long:"trusted-proxies" description:"CIDR ranges that we trust the X-Forwarded-For header from"`
	ValidationKey  string   `env:"VALIDATION_KEY"  long:"validation-key"  required:"true" description:"key used to validate session cookies (32 or 64 bytes)"`
	EncryptionKey  string   `env:"ENCRYPTION_KEY"  long:"encryption-key"  required:"true" description:"key used to encrypt session cookies (32 bytes)"`
}

// ConfigGithub are configurations specifically utilized for interacting with
// Github.
type ConfigGithub struct {
	Token        string `env:"TOKEN"         long:"token"         required:"true" description:"GitHub access token"`
	User         int    `env:"USER"          long:"user"          required:"true" description:"GitHub user ID to use for authentication"`
	ClientID     string `env:"CLIENT_ID"     long:"client-id"     required:"true" description:"GitHub client ID"`
	ClientSecret string `env:"CLIENT_SECRET" long:"client-secret" required:"true" description:"GitHub client secret"`
	SyncOnStart  bool   `env:"SYNC_ON_START" long:"sync-on-start" description:"sync all data from GitHub on startup"`
	CachePath    string `env:"CACHE_PATH"    long:"cache-path"    default:".gitapicache" description:"path to cache directory"`
}

// ConfigWakAPI are configurations specifically utilized for interacting with a
// WakAPI instance (https://github.com/muety/wakapi). This configuration is not
// compatible with Wakatime.
type ConfigWakAPI struct {
	URL    string `env:"URL"     long:"url"     required:"true" description:"wakapi connection url"`
	APIKey string `env:"API_KEY" long:"api-key" required:"true" description:"WakAPI API key"`
}

// ConfigAI are configurations specifically utilized for interacting with
// OpenRouter (https://openrouter.ai).
type ConfigAI struct {
	BaseURL      string `env:"BASE_URL"      long:"base-url"      default:"https://openrouter.ai/api/v1" description:"OpenRouter base URL"`
	APIKey       string `env:"API_KEY"       long:"api-key"       required:"true" description:"OpenRouter API key"`
	DefaultModel string `env:"DEFAULT_MODEL" long:"default-model" description:"default OpenRouter model to use"`
}
