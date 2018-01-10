package model

import (
    "github.com/go-pg/pg/orm"
)

type User struct {
    Id          int     `json:"id"`
    Name        string  `json:"name"`
    Type        string  `json:"type" sql:"-"`
    Relations   []*Relation `json:"-"`
}

func (this *User) AfterQuery(db orm.DB) error {
    this.Type = "user"
    return nil
}
