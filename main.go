// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/flosch/pongo2-addons"
	flags "github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP    HTTPArgs `command:"run" description:"run http server"`
	GenPost GenPost  `command:"gen-post" description:"generate a post"`
}

var (
	cli   Config
	debug = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
)

func main() {
	err := godotenv.Load()
	if err != nil {
		debug.Printf("warn: error loading .env file: %s", err)
	}

	parser := flags.NewParser(&cli, flags.HelpFlag)
	parser.CommandHandler = func(cmd flags.Commander, args []string) error {
		if _, ok := cmd.(*GenPost); ok {
			return cmd.Execute(args)
		}

		return cmd.Execute(args)
	}
	_, err = parser.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(1)
}
