#!/bin/bash
set -e

PGPASSWORD=$POSTGRESQL_PASSWORD psql -v ON_ERROR_STOP=1 --username "$POSTGRESQL_USERNAME" --dbname "$POSTGRESQL_DATABASE" <<-EOSQL
  CREATE TABLE user_tb (
      id serial4 NOT NULL,
      user_id varchar NOT NULL,
      name varchar NOT NULL,
      email varchar,
      created_time timestamp NOT NULL ,
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
    user_id NOT NULL,
    contents text NOT NULL,
    like int4,
    created_time timestamp NOT NULL,
    CONSTRAINT contents_pkey PRIMARY KEY (id)
  );
  ALTER SEQUENCE posting_id_seq restart with 1;

  CREATE TABLE reply (
    id serial4 NOT NULL,
    posting_id serial4 NOT NULL,
    contents varchar NOT NULL,
    created_time timestamp NOT NULL,
    CONSTRAINT contents_pkey PRIMARY KEY (id)
  );
  ALTER SEQUENCE posting_reply_id_seq restart with 1;
EOSQL