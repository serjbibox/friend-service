DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    role VARCHAR(255) NOT NULL,   
    phone VARCHAR(255) NOT NULL UNIQUE, 
    password VARCHAR(255) NOT NULL,    
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    age INTEGER,
    sex VARCHAR(255),    
    icon VARCHAR(255)
);

DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES users(id),
    psy_id INTEGER NOT NULL REFERENCES users(id),
    date TIMESTAMP NOT NULL,
    type VARCHAR(255) NOT NULL,
    session_notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS psy_data;
CREATE TABLE psy_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    intro TEXT NOT NULL,
    video_url VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS client_data;
CREATE TABLE client_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    problem TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);




INSERT INTO users(role, phone, password)
VALUES (
    'client',
    '996550269716',
    '123456'
),
(
    'psy',
    '996550269717',
    '123456'
),
(
    'admin',
    '996550269718',
    '123456'
);