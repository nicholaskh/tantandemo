package model

import (
    "github.com/go-pg/pg/orm"
)

type Relation struct {
    UserId      int     `json:"-" sql:",pk"`
    OtherUserId int     `json:"user_id" sql:",pk"`
    State       string  `json:"state"`
    Type        string  `json:"type" sql:"-"`
}

func (this *Relation) AfterQuery(db orm.DB) error {
    this.Type = "relationship"
    return nil
}

func (this *Relation) AfterInsert(db orm.DB) error {
    this.Type = "relationship"
    return nil
}
