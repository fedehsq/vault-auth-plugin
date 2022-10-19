/*  
    psql -d postgres -U fedeveloper 
    \c myDb;
*/
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