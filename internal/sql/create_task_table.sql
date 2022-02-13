CREATE TABLE task (
    id serial4 NOT NULL,
    run_time int4 NOT NULL DEFAULT 0,
    email text NOT NULL DEFAULT ''::text,
    "name" text NOT NULL DEFAULT ''::text,
    CONSTRAINT task_pkey PRIMARY KEY (id)
);