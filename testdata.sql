DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS debitors;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS payments;
CREATE TABLE users (id INTEGER, debitor_id INTEGER);
CREATE TABLE debitors (id INTEGER, name VARCHAR(42));
CREATE TABLE accounts (id INTEGER, name VARCHAR(42), debitor_id INTEGER);
CREATE TABLE payments (id INTEGER, account_id INTEGER, amount FLOAT, balance FLOAT, paid_at TEXT);

INSERT INTO users (id, debitor_id) VALUES (1, 1);
INSERT INTO debitors (id, name) VALUES (1, "Chris Marshall");
INSERT INTO accounts (id, name, debitor_id) VALUES (1, "Citibank", 1), 
                                                   (2, "Bank of America", 1),
                                                   (3, "Line of Credit", 1),
                                                   (4, "Sallie Mae", 1);

INSERT INTO payments (id, account_id, amount, balance, paid_at) VALUES
    (1, 1, 200, 1000.57, datetime("2015-01-01 01:00:00")),
    (2, 2, 400,  500,    datetime("2015-01-01 01:00:00")),
    (3, 3, 260, 4568.27, datetime("2015-01-01 01:00:00")),
    (4, 4,  20,  290.16, datetime("2015-01-01 01:00:00")),
    (5, 1,   5,   50.85, datetime("2015-01-12 01:00:00"));

