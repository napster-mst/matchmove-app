package ent

import "time"

type Roles struct {
	tableName struct{} `pg:"public.roles"`
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
