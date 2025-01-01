-- migrate:up
CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE diaries (
    id VARCHAR(26) PRIMARY KEY,
    user_id VARCHAR(26) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    visibility VARCHAR(10) NOT NULL DEFAULT 'public',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE diary_tags (
    id VARCHAR(26) PRIMARY KEY,
    diary_id VARCHAR(26) NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (diary_id) REFERENCES diaries(id) ON DELETE CASCADE,
    UNIQUE KEY uk_diary_tag (diary_id, name)
);

-- インデックスの追加
CREATE INDEX idx_diaries_user_id ON diaries(user_id);
CREATE INDEX idx_diaries_visibility ON diaries(visibility);
CREATE INDEX idx_diary_tags_name ON diary_tags(name);

-- migrate:down
DROP TABLE IF EXISTS diary_tags;
DROP TABLE IF EXISTS diaries;
DROP TABLE IF EXISTS users;
