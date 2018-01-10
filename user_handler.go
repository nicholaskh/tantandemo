package main

import (
    "log"
    "net/http"

    "github.com/nicholaskh/tantandemo/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []*model.User
    db := GetDb()
    err := db.Model(&users).Select()
    if err != nil {
        log.Print(err)
        JsonResp(w, http.StatusInternalServerError, err.Error())
        return
    }
    JsonResp(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    code, err := ProcJsonBody(r, &user)
    if err != nil {
        log.Print(err)
        JsonResp(w, code, err.Error())
        return
    }

    db := GetDb()
    _, err = db.Model(&user).Returning("id").Insert()
    if err != nil {
        log.Print(err)
        JsonResp(w, http.StatusInternalServerError, err.Error())
        return
    }
    JsonResp(w, http.StatusCreated, user)
}
