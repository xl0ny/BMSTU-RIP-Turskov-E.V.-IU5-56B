CREATE TABLE IF NOT EXISTS users (
    id           BIGSERIAL PRIMARY KEY,
    login        VARCHAR(32) UNIQUE NOT NULL,
    password     VARCHAR(128)      NOT NULL,
    is_moderator BOOLEAN           NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS criteria (
    id          BIGSERIAL     PRIMARY KEY,
    code        VARCHAR(8)    NOT NULL,
    name        VARCHAR(120)  NOT NULL,
    indicator   VARCHAR(80)   NOT NULL,
    duration    VARCHAR(60)   NOT NULL,
    home_visit  BOOLEAN       NOT NULL DEFAULT FALSE,
    image_url   TEXT          NOT NULL,      -- URL в MinIO
    description TEXT          NOT NULL,
    is_active   BOOLEAN       NOT NULL DEFAULT TRUE
    );

CREATE TABLE IF NOT EXISTS orders (
    id              BIGSERIAL PRIMARY KEY,
    status          VARCHAR(16) NOT NULL, -- draft|deleted|formed|finished|rejected
    created_at      TIMESTAMP   NOT NULL DEFAULT NOW(),
    creator_id      BIGINT      NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    formed_at       TIMESTAMP   NULL,
    finished_at     TIMESTAMP   NULL,
    moderator_id    BIGINT      NULL REFERENCES users(id) ON DELETE RESTRICT,
    computed_result VARCHAR(120) NULL
    );

CREATE TABLE IF NOT EXISTS order_items (
    id            BIGSERIAL PRIMARY KEY,
    order_id      BIGINT NOT NULL REFERENCES orders(id)   ON DELETE RESTRICT,
    criterion_id  BIGINT NOT NULL REFERENCES criteria(id) ON DELETE RESTRICT,
    position      INTEGER      NOT NULL DEFAULT 0,
    value_num     NUMERIC(10,3) NULL,
    value_indicator     BOOLEAN NULL,
    UNIQUE (order_id, criterion_id)
    );