// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package models

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
	"github.com/jessevdk/go-flags"
)

type Flags struct {
	Debug    bool           `long:"debug" env:"DEBUG" description:"enable debugging (pprof endpoints) and disables template caching"`
	Log      ConfigLogger   `group:"Logging Options" namespace:"log" env-namespace:"LOG"`
	HTTP     ConfigHTTP     `group:"HTTP Server options" namespace:"http" env-namespace:"HTTP"`
	Database ConfigDatabase `group:"Database options" namespace:"database" env-namespace:"DATABASE"`
	Github   ConfigGithub   `group:"Github options" namespace:"github" env-namespace:"GITHUB"`
}

func (cli *Flags) Parse() (args []string) {
	parser := flags.NewParser(cli, flags.HelpFlag|flags.PassDoubleDash)
	parser.NamespaceDelimiter = "."
	parser.EnvNamespaceDelimiter = "_"

	var err error

	args, err = parser.Parse()
	if err != nil {
		if FlagErr, ok := err.(*flags.Error); ok && FlagErr.Type == flags.ErrHelp {
			os.Exit(0)
		}

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return args
}

// ConfigDatabase holds the database configuration.
type ConfigDatabase struct {
	URL string `env:"URL" long:"url" required:"true" description:"database connection url"`
}

// ConfigLogger are the flags that define how log entries are processed/returned.
type ConfigLogger struct {
	Quiet  bool   `env:"QUIET"  long:"quiet"  description:"disable logging to stdout (also: see levels)"`
	Level  string `env:"LEVEL"  long:"level"  default:"info" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal"  description:"logging level"`
	JSON   bool   `env:"JSON"   long:"json"   description:"output logs in JSON format"`
	Pretty bool   `env:"PRETTY" long:"pretty" description:"output logs in a pretty colored format (cannot be easily parsed)"`
}

// ParseConfigLogger parses ConfigLogger and adjusts the structured logger as
// necessary.
func (c ConfigLogger) Parse(debug bool) log.Interface {
	logger := &log.Logger{}

	if debug {
		logger.Level = log.DebugLevel
	} else {
		logger.Level = log.MustParseLevel(c.Level)
	}

	switch {
	case c.Quiet:
		logger.Handler = discard.New()
	case c.JSON:
		logger.Handler = json.New(os.Stdout)
	case c.Pretty:
		logger.Handler = text.New(os.Stdout)
	default:
		logger.Handler = logfmt.New(os.Stdout)
	}

	// Set global options as well, just in case.
	log.SetLevel(logger.Level)
	log.SetHandler(logger.Handler)

	return logger
}

// ConfigHTTP are configurations specifically utilized by the HTTP service.
type ConfigHTTP struct {
	BaseURL        string   `env:"BASE_URL"        long:"base-url"        default:"http://localhost:8080" description:"base url for the HTTP server"`
	BindAddr       string   `env:"BIND_ADDR"       long:"bind-addr"       default:":8080" required:"true" description:"ip:port pair to bind to"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" long:"trusted-proxies" description:"CIDR ranges that we trust the X-Forwarded-For header from"`
	SessionKey     string   `env:"SESSION_KEY"     long:"session-key"     required:"true" description:"key used to validate session cookies (32 or 64 bytes)"`
}

// ConfigGithub are configurations specifically utilized for interacting with
// Github.
type ConfigGithub struct {
	Token        string `env:"TOKEN"         long:"token"         required:"true" description:"GitHub access token"`
	User         int    `env:"USER"          long:"user"          required:"true" description:"GitHub user ID to use for authentication"`
	ClientID     string `env:"CLIENT_ID"     long:"client-id"     required:"true" description:"GitHub client ID"`
	ClientSecret string `env:"CLIENT_SECRET" long:"client-secret" required:"true" description:"GitHub client secret"`
}
