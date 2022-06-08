package ent

import "time"

type Token struct {
	tableName struct{} `pg:"public.tokens"`
	ID        int
	Token     string
	StatusId  int
	UserId    int
	Expiry    time.Time
}
