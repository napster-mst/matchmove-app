package model

const (
	//Sequence
	CreateSequenceRole = `CREATE SEQUENCE IF NOT EXISTS public.roles_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1`

	// Tables
	CreateRoleTable = `CREATE TABLE IF NOT EXISTS public.roles
	(
		id integer NOT NULL DEFAULT nextval('roles_id_seq'::regclass),
		name text NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone,
		CONSTRAINT roles_pkey PRIMARY KEY (id),
		CONSTRAINT roles_name_key UNIQUE (name)
	)`
)
