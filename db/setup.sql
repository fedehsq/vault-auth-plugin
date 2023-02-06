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

CREATE TABLE "remote_hosts" (
    id SERIAL PRIMARY KEY,
    ip TEXT
);

CREATE TABLE "remote_host_users" (
    id SERIAL PRIMARY KEY,
    remote_host_id INTEGER REFERENCES remote_hosts(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE "logs" (
    id SERIAL PRIMARY KEY,
    time TEXT,
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