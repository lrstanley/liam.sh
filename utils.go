package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var debug *log.Logger

func initLogger() {
	if conf.DebugLog != "" {
		f, err := os.OpenFile(conf.DebugLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		checkErr(err)

		debug = log.New(io.MultiWriter(os.Stdout, f), "", log.Lshortfile|log.LstdFlags)
		return
	}

	debug = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	debug.Print("initialized logger")
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func httpPanic(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
