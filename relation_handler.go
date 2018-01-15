package main

import (
    "errors"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/go-pg/pg"

    "github.com/nicholaskh/tantandemo/model"
)

func GetUserRelations(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var userId int
    var err error
    if userId, err = strconv.Atoi(vars["user_id"]); err != nil {
        log.Print(err)
        JsonResp(w, http.StatusUnprocessableEntity, err.Error())
        return
    }
    db := GetDb()
    user := model.User{Id: userId}
    err = db.Model(&user).
        Column("user.*", "Relations").
        First()
    if err != nil {
        log.Print(err)
        JsonResp(w, http.StatusInternalServerError, err.Error())
        return
    }
    JsonResp(w, http.StatusOK, user.Relations)
}

func SetUserRelation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var userId int
    var otherUserId int
    var err error
    if userId, err = strconv.Atoi(vars["user_id"]); err != nil {
        log.Print(err)
        JsonResp(w, http.StatusUnprocessableEntity, err.Error())
        return
    }
    if otherUserId, err = strconv.Atoi(vars["other_user_id"]); err != nil {
        log.Print(err)
        JsonResp(w, http.StatusUnprocessableEntity, err.Error())
        return
    }
    if userId == otherUserId {
        err := errors.New("Could not like/dislike self")
        log.Print(err)
        JsonResp(w, http.StatusUnprocessableEntity, err.Error())
        return
    }

    var relation model.Relation
    code, err := ProcJsonBody(r, &relation)
    if err != nil {
        log.Print(err)
        JsonResp(w, code, err.Error())
        return
    }
    if relation.State != "liked" && relation.State != "disliked" {
        err := errors.New("State should be liked or disliked")
        log.Print(err)
        JsonResp(w, http.StatusUnprocessableEntity, err.Error())
        return
    }

    db := GetDb()
    oppRelation := model.Relation{UserId: otherUserId, OtherUserId: userId}
    err = db.Select(&oppRelation)
    if err != nil && err != pg.ErrNoRows {
        log.Print(err)
        JsonResp(w, http.StatusInternalServerError, err.Error())
        return
    }

    relation.UserId = userId
    relation.OtherUserId = otherUserId
    if relation.State == "liked" && (oppRelation.State == "liked" || oppRelation.State == "matched") {
        relation.State = "matched"
        if oppRelation.State == "liked" {
            oppRelation.State = "matched"
            db.Model(&oppRelation).Update()
        }
    } else if relation.State == "disliked" && oppRelation.State == "matched" {
        oppRelation.State = "liked"
        db.Model(&oppRelation).Update()
    }

    _, err = db.Model(&relation).
        OnConflict("(user_id, other_user_id) DO UPDATE").
        Set("state = EXCLUDED.state").
        Insert()
    if err != nil {
        log.Print(err)
        JsonResp(w, http.StatusInternalServerError, err.Error())
        return
    }

    JsonResp(w, http.StatusCreated, relation)
}
