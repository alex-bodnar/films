-- +migrate Up
CREATE TABLE IF NOT EXISTS responce_time_log (
    id          SERIAL PRIMARY KEY,
    request     TEXT UNIQUE NOT NULL,
    time_db     BIGINT,
    time_redis  BIGINT,
    time_memory BIGINT
);

-- +migrate Down
DROP TABLE IF EXISTS responce_time_log;
