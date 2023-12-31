CREATE TABLE IF NOT EXISTS Users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    login VARCHAR(40),
    password VARCHAR(120),
    email VARCHAR(100),
    is_banned BOOLEAN DEFAULT false,
    banned_until DATE,
    ban_count INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Roles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) PRIMARY KEY NOT NULL
);

CREATE TABLE IF NOT EXISTS UserRoles (
    user_id BIGSERIAL REFERENCES Users(id) ON DELETE CASCADE,
    role_id SERIAL REFERENCES Roles(id) ON DELETE SET NULL
);
