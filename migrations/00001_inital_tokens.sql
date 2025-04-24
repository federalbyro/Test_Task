CREATE TABLE tokens (
  access_token_id text PRIMARY KEY,
  user_id         text,
  ip_address      text,
  refresh_hash    text,
  flag_reuse      boolean
);

CREATE UNIQUE INDEX flag_index ON tokens (access_token_id) WHERE flag_reuse = false;

