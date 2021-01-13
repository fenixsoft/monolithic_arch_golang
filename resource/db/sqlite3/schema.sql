DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS specifications;
DROP TABLE IF EXISTS advertisements;
DROP TABLE IF EXISTS stockpiles;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS payments;

CREATE TABLE accounts
(
    id        INTEGER IDENTITY PRIMARY KEY,
    username  VARCHAR(50),
    password  VARCHAR(100),
    name      VARCHAR(50),
    avatar    VARCHAR(100),
    telephone VARCHAR(20),
    email     VARCHAR(100),
    location  VARCHAR(100)
);
CREATE UNIQUE INDEX accounts_user ON accounts (username);
CREATE UNIQUE INDEX accounts_telephone ON accounts (telephone);
CREATE UNIQUE INDEX accounts_email ON accounts (email);

CREATE TABLE wallets
(
    id         INTEGER IDENTITY PRIMARY KEY,
    money      DECIMAL,
    account_id INTEGER,
    CONSTRAINT fk_wallets_accounts FOREIGN KEY (account_id) REFERENCES accounts (id)
);


CREATE TABLE products
(
    id          INTEGER IDENTITY PRIMARY KEY,
    title       VARCHAR(50),
    price       DECIMAL,
    rate        FLOAT,
    description VARCHAR(8000),
    cover       VARCHAR(100),
    detail      VARCHAR(100)
);
CREATE INDEX products_title ON products (title);

CREATE TABLE stockpiles
(
    id         INTEGER IDENTITY PRIMARY KEY,
    amount     INTEGER,
    frozen     INTEGER,
    product_id INTEGER,
    CONSTRAINT fk_stockpiles_products FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE specifications
(
    id         INTEGER IDENTITY PRIMARY KEY,
    item       VARCHAR(50),
    value      VARCHAR(100),
    product_id INTEGER,
    CONSTRAINT fk_specifications_products FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE advertisements
(
    id         INTEGER IDENTITY PRIMARY KEY,
    image      VARCHAR(100),
    product_id INTEGER,
    CONSTRAINT fk_advertisements_products FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE payments
(
    id           INTEGER IDENTITY PRIMARY KEY,
    pay_id       VARCHAR(100),
    create_time  DATETIME,
    total_price  DECIMAL,
    expires      INTEGER NOT NULL,
    payments_link VARCHAR(300),
    pay_state    VARCHAR(20)
);
