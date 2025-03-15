-- +goose Up
CREATE TABLE IF NOT EXISTS public.tasks
(
    id serial,
    title text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    status text COLLATE pg_catalog."default" DEFAULT 'new'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT tasks_pkey PRIMARY KEY (id),
    CONSTRAINT check_status CHECK (status = 'new'::text OR status = 'in_progress'::text OR status = 'done'::text) NOT VALID
);
-- +goose Down
DROP TABLE public.tasks;