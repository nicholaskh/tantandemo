package main

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
)

func ProcJsonBody(r *http.Request, obj interface{}) (int, error) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        return http.StatusInternalServerError, err
    }
    if err := r.Body.Close(); err != nil {
        return http.StatusInternalServerError, err
    }

    if err := json.Unmarshal(body, obj); err != nil {
        return http.StatusUnprocessableEntity, err
    }
    return 0, nil
}
