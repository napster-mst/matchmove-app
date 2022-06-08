package model

const (
	//Sequence
	CreateSequenceTokens = `CREATE SEQUENCE IF NOT EXISTS public.token_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1`

	// Tables
	CreateTokensTable = `CREATE TABLE IF NOT EXISTS public.tokens
	(
		id integer NOT NULL DEFAULT nextval('token_id_seq'::regclass),
		token text COLLATE pg_catalog."default" NOT NULL,
		status_id integer NOT NULL,
		user_id integer NOT NULL,
		expiry timestamp with time zone,
		CONSTRAINT tokens_pkey PRIMARY KEY (id),
		CONSTRAINT token_name_key UNIQUE (token),
		CONSTRAINT status_token_id_fkey FOREIGN KEY (status_id)
			REFERENCES public.status (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION,
		CONSTRAINT users_token_id_fkey FOREIGN KEY (user_id)
			REFERENCES public.users (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
	)`
)
