CREATE TABLE task (
    id serial4 NOT NULL,
    run_time int4 NOT NULL DEFAULT 0,
    user_id int4 NULL,
    CONSTRAINT task_pkey PRIMARY KEY (id),
    CONSTRAINT task_to_users_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);