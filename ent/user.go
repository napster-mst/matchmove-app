package ent

type User struct {
	tableName struct{} `pg:"public.users"`
	ID        int
	Username  string
	Password  string
	RoleID    int
}
