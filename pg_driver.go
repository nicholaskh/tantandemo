package main

import (
    "github.com/go-pg/pg"
)

var db *pg.DB

func GetDb() *pg.DB {
    if db != nil {
        return db
    }
    db = pg.Connect(&pg.Options{
        User: "postgres",
        Password: "123456",
        Database: "test",
    })

    return db
}
