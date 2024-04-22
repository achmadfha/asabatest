CREATE TYPE transaction_type AS ENUM ('IN', 'OUT');
CREATE TYPE roles AS ENUM ('ADMIN', 'USER');

CREATE TABLE users
(
    users_id   UUID PRIMARY KEY,
    email      VARCHAR(100) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    roles      roles       NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items
(
    items_id     UUID PRIMARY KEY,
    code         VARCHAR(10) NOT NULL UNIQUE,
    name         VARCHAR(50) NOT NULL,
    amount       INT         NOT NULL,
    description  TEXT,
    statusActive BOOLEAN   DEFAULT TRUE,
    isDeleted    BOOLEAN   DEFAULT FALSE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transaction_items
(
    transaction_id   UUID PRIMARY KEY,
    items_code       VARCHAR(10)      NOT NULL,
    transaction_type transaction_type NOT NULL,
    quantity         INT              NOT NULL,
    isDeleted    BOOLEAN   DEFAULT FALSE,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);