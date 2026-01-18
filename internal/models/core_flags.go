// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

type Flags struct {
	HTTP     ConfigHTTP     `embed:"" group:"HTTP Server options" prefix:"http."     envprefix:"HTTP_"`
	Database ConfigDatabase `embed:"" group:"Database options"    prefix:"database." envprefix:"DATABASE_"`
	Github   ConfigGithub   `embed:"" group:"Github options"      prefix:"github."   envprefix:"GITHUB_"`
	WakAPI   ConfigWakAPI   `embed:"" group:"WakAPI options"      prefix:"wakapi."   envprefix:"WAKAPI_"`
	AI       ConfigAI       `embed:"" group:"AI options"          prefix:"ai."       envprefix:"AI_"`
}

// ConfigDatabase holds the database configuration.
type ConfigDatabase struct {
	URL string `env:"URL" default:"file:local.db" help:"database connection url"`
}

// ConfigHTTP are configurations specifically utilized by the HTTP service.
type ConfigHTTP struct {
	BaseURL        string   `env:"BASE_URL"        name:"base-url"        default:"http://localhost:8080" help:"base url for the HTTP server"`
	BindAddr       string   `env:"BIND_ADDR"       name:"bind-addr"       default:":8080" required:"" help:"ip:port pair to bind to"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" name:"trusted-proxies" help:"CIDR ranges that we trust the X-Forwarded-For header from"`
	ValidationKey  string   `env:"VALIDATION_KEY"  name:"validation-key"  required:"" help:"key used to validate session cookies (32 or 64 bytes)"`
	EncryptionKey  string   `env:"ENCRYPTION_KEY"  name:"encryption-key"  required:"" help:"key used to encrypt session cookies (32 bytes)"`
}

// ConfigGithub are configurations specifically utilized for interacting with
// Github.
type ConfigGithub struct {
	Token        string `env:"TOKEN"         name:"token"         required:"" help:"GitHub access token"`
	User         int    `env:"USER"          name:"user"          required:"" help:"GitHub user ID to use for authentication"`
	ClientID     string `env:"CLIENT_ID"     name:"client-id"     required:"" help:"GitHub client ID"`
	ClientSecret string `env:"CLIENT_SECRET" name:"client-secret" required:"" help:"GitHub client secret"`
	SyncOnStart  bool   `env:"SYNC_ON_START" name:"sync-on-start" help:"sync all data from GitHub on startup"`
	CachePath    string `env:"CACHE_PATH"    name:"cache-path"    default:".gitapicache" help:"path to cache directory"`
}

// ConfigWakAPI are configurations specifically utilized for interacting with a
// WakAPI instance (https://github.com/muety/wakapi). This configuration is not
// compatible with Wakatime.
type ConfigWakAPI struct {
	URL    string `env:"URL"     name:"url"     required:"" help:"wakapi connection url"`
	APIKey string `env:"API_KEY" name:"api-key" required:"" help:"WakAPI API key"`
}

// ConfigAI are configurations specifically utilized for interacting with
// OpenRouter (https://openrouter.ai).
type ConfigAI struct {
	BaseURL      string `env:"BASE_URL"      name:"base-url"      default:"https://openrouter.ai/api/v1" help:"OpenRouter base URL"`
	APIKey       string `env:"API_KEY"       name:"api-key"       required:"" help:"OpenRouter API key"`
	DefaultModel string `env:"DEFAULT_MODEL" name:"default-model" help:"default OpenRouter model to use"`
}
