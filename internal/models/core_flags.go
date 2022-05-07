// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

type Flags struct {
	HTTP     ConfigHTTP     `group:"HTTP Server options" namespace:"http" env-namespace:"HTTP"`
	Database ConfigDatabase `group:"Database options" namespace:"database" env-namespace:"DATABASE"`
	Github   ConfigGithub   `group:"Github options" namespace:"github" env-namespace:"GITHUB"`
}

// ConfigDatabase holds the database configuration.
type ConfigDatabase struct {
	URL string `env:"URL" long:"url" required:"true" description:"database connection url"`
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
}
