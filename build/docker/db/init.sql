CREATE TABLE IF NOT EXISTS sections (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS cert_area (
    id SERIAL NOT NULL PRIMARY KEY,
    section_id INTEGER REFERENCES sections (id),
    name VARCHAR(512) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS tests (
    id SERIAL NOT NULL PRIMARY KEY,
    cert_area_id INTEGER REFERENCES cert_area (id),
    name VARCHAR(512) NOT NULL
);
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL NOT NULL PRIMARY KEY,
    test_id INTEGER REFERENCES tests (id),
    link VARCHAR(128) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL NOT NULL PRIMARY KEY,
    test_id INTEGER REFERENCES tests (id),
    question VARCHAR(512) NOT NULL UNIQUE,
    answer VARCHAR(512) NOT NULL
);
CREATE TABLE IF NOT EXISTS options (
    id SERIAL NOT NULL PRIMARY KEY,
    question_id INTEGER REFERENCES tasks (id) ON DELETE CASCADE,
    answer_option VARCHAR(512) NOT NULL,
    CONSTRAINT unique_options UNIQUE (question_id, answer_option)
);