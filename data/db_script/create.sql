CREATE TABLE user_tb (
    id serial4 NOT NULL,
    name varchar NOT NULL,
    email varchar,
    created_time timestamp NOT NULL,
  CONSTRAINT user_pkey PRIMARY KEY (id)
);
ALTER SEQUENCE user_tb_id_seq restart with 1;

CREATE TABLE device (
  id serial4 NOT NULL,
  device_token varchar NOT NULL,
  user_id int4 NOT NULL,
  last_logged_in_at timestamp NOT NULL,
  CONSTRAINT device_pkey PRIMARY KEY (id)
);
ALTER SEQUENCE device_id_seq restart with 1;

CREATE TABLE posting (
  id serial4 NOT NULL,
  user_id int4 NOT NULL,
  contents text NOT NULL,
  likes int4,
  created_time timestamp NOT NULL,
  CONSTRAINT posting_pkey PRIMARY KEY (id)
);
ALTER SEQUENCE posting_id_seq restart with 1;

CREATE TABLE reply (
  id serial4 NOT NULL,
  posting_id serial4 NOT NULL,
  contents varchar NOT NULL,
  created_time timestamp NOT NULL,
  CONSTRAINT reply_pkey PRIMARY KEY (id)
);
ALTER SEQUENCE reply_id_seq restart with 1;

CREATE TABLE follower (
    id serial4 NOT NULL,
    user_id integer NOT NULL,
    followers integer[],
    CONSTRAINT follower_pkey PRIMARY KEY (id)
);
ALTER SEQUENCE follower_id_seq restart with 1;
