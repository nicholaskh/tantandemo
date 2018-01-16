package main

import (
    "github.com/go-pg/pg"
    "sync"
)

var db *pg.DB
var once sync.Once

func GetDb() *pg.DB {
    once.Do(func() {
        db = pg.Connect(&pg.Options{
            User: "postgres",
            Password: "123456",
            Database: "test",
        })
    })
    return db
}
