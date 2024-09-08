-- +goose Up
-- +goose StatementBegin

CREATE TABLE public.users (
	id serial4 NOT NULL,
	username varchar NOT NULL,
	password_hash varchar NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

INSERT INTO public.users (username, password_hash) VALUES ('test', 'eb8d56f1a3113446c89268f994d5f0bf6c4118ea');

CREATE TABLE public.images (
	id serial4 NOT NULL,
	user_id int4 NOT NULL,
	image_path varchar NOT NULL,
	image_url varchar NOT NULL,
	CONSTRAINT images_pk PRIMARY KEY (id)
);

ALTER TABLE public.images ADD CONSTRAINT images_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE images;
DROP TABLE users;
-- +goose StatementEnd
