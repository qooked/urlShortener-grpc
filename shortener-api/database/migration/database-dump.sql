CREATE TABLE IF NOT EXISTS url (
    id SERIAL PRIMARY KEY,
    shortened_url TEXT NOT NULL,
    original_url TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS url_alias_key ON url (shortened_url);
CREATE INDEX IF NOT EXISTS idx_alias ON url (shortened_url);
