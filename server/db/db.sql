/* psql -d postgres -U fedeveloper */
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    Username TEXT,
    Password TEXT
);