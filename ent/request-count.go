package ent

type RequestCounts struct {
	tableName struct{} `pg:"public.request_counts"`
	ID        int
	TokenID   int
	Count     int
}
