DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    role VARCHAR(255) NOT NULL,    
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255),    
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    video_url VARCHAR(255),
    tag VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES users(id),
    psychologist_id INTEGER NOT NULL REFERENCES users(id),
    date TIMESTAMP NOT NULL,
    type VARCHAR(255) NOT NULL,
    session_notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS anketa;
CREATE TABLE anketa (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES users(id),
    age INTEGER NOT NULL,
    sex VARCHAR(255) NOT NULL,
    problem TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);