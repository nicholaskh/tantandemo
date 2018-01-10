package main

import (
    "log"
    "net/http"

    "github.com/nicholaskh/tantandemo/config"
    "github.com/nicholaskh/golib/server"
)

var cf *config.ConfigTest

func init() {
    parseFlags()
}

func main() {
    conf := server.LoadConfig(options.configFile)
    cf = new(config.ConfigTest)
    cf.LoadConfig(conf)

    router := NewRouter()

    log.Fatal(http.ListenAndServe(cf.ListenAddr, router))
}
