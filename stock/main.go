package main

import (
	"flag"
	"log"

	"github.com/amosehiguese/stock/app/config"
	"github.com/amosehiguese/stock/app/server"
)

var addr = flag.String("addr", "127.0.0.1:8080", "server addr to serve on.")
var readTimeout = flag.Int64("read-timeout", 10, "maximum duration for reading the entire request")
var writeTimeout = flag.Int64("write-timeout", 600, "maximum amount of time to wait for the next request")
var certFile = flag.String("cert-file", "", "path to cert-file")
var KeyFile = flag.String("key-file", "", "path to key-file")

func main() {
	flag.Parse()

	config := config.Config{
		Address: *addr,
		ReadTimeout: *readTimeout,
		WriteTimeout: *writeTimeout,
		CertFile: *certFile,
		KeyFile: *KeyFile,
	}

	srv := server.NewHttpServer(config)
	if config.CertFile == "" && config.KeyFile == "" {
		if err :=srv.ListenAndServe(); err != nil {
			log.Fatal("Error starting server ->", err)
		}
	} else {
		if err := srv.ListenAndServeTLS(config.CertFile, config.KeyFile); err != nil {
			log.Fatalln("Error starting server over TLS ->", err)
		}
	}
}
