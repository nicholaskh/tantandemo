package main

import (
    "github.com/go-pg/pg"
    "sync"
)

var db *pg.DB
var mutex *sync.Mutex

func init() {
    mutex = new(sync.Mutex)
}

func GetDb() *pg.DB {
    if db != nil {
        return db
    }
    mutex.Lock()
    defer mutex.Unlock()
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
