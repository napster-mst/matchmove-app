package ent

type Status struct {
	tableName struct{} `pg:"public.status"`
	ID        int
	Name      string
}
