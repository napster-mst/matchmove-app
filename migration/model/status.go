package model

const (
	//Sequence
	CreateSequenceStatus = `CREATE SEQUENCE IF NOT EXISTS public.status_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1`

	// Tables
	CreateStatusTable = `CREATE TABLE IF NOT EXISTS public.status
	(
		id integer NOT NULL DEFAULT nextval('status_id_seq'::regclass),
		name text NOT NULL,
		CONSTRAINT status_pkey PRIMARY KEY (id),
		CONSTRAINT status_name_key UNIQUE (name)
	)`
)
