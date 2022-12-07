CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    Username TEXT,
    Password TEXT
);

CREATE TABLE "admins" (
    id SERIAL PRIMARY KEY,
    Username TEXT,
    Password TEXT
);

CREATE TABLE "logs" (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP,
    ip TEXT,
    caller_identity TEXT,
    method TEXT,
    route TEXT,
    body TEXT
);

/* Add a user to the database */
INSERT INTO users (Username, Password) VALUES ('elliot', 'mrrobot');

/* Add an admin to the database */
INSERT INTO admins (Username, Password) VALUES ('admin', 'admin');