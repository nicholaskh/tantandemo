package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func JsonResp(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Print(err)
        return
    }
}
