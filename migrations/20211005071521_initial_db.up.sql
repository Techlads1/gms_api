 
-- SEQUENCE: public.regions_id_seq

CREATE SEQUENCE public.regions_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.regions_id_seq
    OWNER TO postgres;

-- CREATE regions
	
CREATE TABLE IF NOT EXISTS public.regions
(
    id integer NOT NULL DEFAULT nextval('regions_id_seq'::regclass),
    name character varying(45) COLLATE pg_catalog."default" NOT NULL,
    code character varying(10) COLLATE pg_catalog."default",
    description character varying(30) COLLATE pg_catalog."default",
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    CONSTRAINT regions_pkey PRIMARY KEY (id)
)


-- SEQUENCE: public.users_id_seq

-- DROP SEQUENCE public.users_id_seq;

CREATE SEQUENCE public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.users_id_seq
    OWNER TO postgres;
    
-- Table: public.users

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    name character varying(45) COLLATE pg_catalog."default" NOT NULL,
    email character varying(90) COLLATE pg_catalog."default" NOT NULL,
    email_verified_at timestamp(0) without time zone,
    password character varying(191) COLLATE pg_catalog."default" NOT NULL,
    login_attempt integer NOT NULL DEFAULT 0,
    user_category_id integer,
    remember_token character varying(100) COLLATE pg_catalog."default",
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_unique UNIQUE (email)
)

TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;