-- Создание таблицы url
CREATE TABLE IF NOT EXISTS url (
    id SERIAL PRIMARY KEY,
    shortened_url TEXT NOT NULL,
    original_url TEXT NOT NULL
);

-- Создание уникального индекса для поля shortened_url
CREATE UNIQUE INDEX IF NOT EXISTS url_alias_key ON url (shortened_url);

-- Создание индекса для поля shortened_url
CREATE INDEX IF NOT EXISTS idx_alias ON url (shortened_url);
