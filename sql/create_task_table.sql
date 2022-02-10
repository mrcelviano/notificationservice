CREATE TABLE task (
                             id serial4 NOT NULL,
                             run_time int4 NOT NULL DEFAULT 0,
                             "user" json NOT NULL DEFAULT '{}'::json,
                             CONSTRAINT task_pkey PRIMARY KEY (id)
);