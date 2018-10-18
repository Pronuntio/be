CREATE ROLE pr_user WITH LOGIN PASSWORD 'pr_user_pass';

CREATE SCHEMA pronuntio;

GRANT ALL PRIVILEGES ON DATABASE pr_main TO pr_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA pronuntio TO pr_user;

CREATE SEQUENCE pronuntio.w_pk_seq START WITH 100000000;

CREATE TYPE pronuntio.word_status_type AS ENUM ('private', 'public');

CREATE TABLE pronuntio.words (
    id INTEGER UNIQUE PRIMARY KEY DEFAULT nextval('pronuntio.w_pk_seq'::regclass) NOT NULL,
    text_original TEXT NOT NULL,
    text_english TEXT NOT NULL,
    status pronuntio.word_status_type NOT NULL,
    filename VARCHAR(250) NOT NULL
);

