package model

const (
	// Tables
	CreateReqCountTable = `CREATE TABLE IF NOT EXISTS public.request_counts
	(
		 id serial primary key,
		token_id integer NOT NULL,
		count integer NOT NULL,
		CONSTRAINT token_request_id_fkey FOREIGN KEY (token_id)
			REFERENCES public.tokens (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
	)`
)
