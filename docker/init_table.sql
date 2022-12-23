create extension if not exists pg_stat_statements cascade;

CREATE TABLE IF NOT EXISTS public.test
(
    id integer NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_timestamp timestamp without time zone DEFAULT now(),
    CONSTRAINT test_pkey PRIMARY KEY (id)
);
