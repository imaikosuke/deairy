-- migrate:up
INSERT INTO users (id, name, email) VALUES
('01HN5GZPV6XF3RS8K9AZ4XD1GJ', 'テストユーザー1', 'test1@example.com'),
('01HN5GZPV6XF3RS8K9AZ4XD1GK', 'テストユーザー2', 'test2@example.com');

INSERT INTO diaries (id, user_id, title, content, visibility) VALUES
('01HN5GZPV6XF3RS8K9AZ4XD1GL', '01HN5GZPV6XF3RS8K9AZ4XD1GJ', '最初の日記', 'これは最初の日記です。', 'public'),
('01HN5GZPV6XF3RS8K9AZ4XD1GM', '01HN5GZPV6XF3RS8K9AZ4XD1GJ', '非公開の日記', 'これは非公開の日記です。', 'private');

INSERT INTO diary_tags (id, diary_id, name) VALUES
('01HN5GZPV6XF3RS8K9AZ4XD1GN', '01HN5GZPV6XF3RS8K9AZ4XD1GL', '日常'),
('01HN5GZPV6XF3RS8K9AZ4XD1GP', '01HN5GZPV6XF3RS8K9AZ4XD1GL', '思考');

-- migrate:down
DELETE FROM diary_tags;
DELETE FROM diaries;
DELETE FROM users;
