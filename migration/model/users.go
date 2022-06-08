package model

const (
	//Sequence
	CreateSequenceUser = `CREATE SEQUENCE IF NOT EXISTS public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1`

	// Tables
	CreateUsersTable = `CREATE TABLE IF NOT EXISTS public.users
	(
		id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
		username text NOT NULL,
		password text NOT NULL,
		role_id integer NOT NULL,
		CONSTRAINT users_pkey PRIMARY KEY (id),
		CONSTRAINT user_name_key UNIQUE (username),
		CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id)
			REFERENCES public.roles (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
	)`
)
