-- +goose Up
-- +goose StatementBegin

-- Администратор по умолчанию (пароль: admin123)
-- bcrypt hash сгенерирован с cost=12
INSERT INTO users (id, username, password_hash, full_name, is_admin)
VALUES (
    'a3f1e2d4-7b6c-4e8f-9a0b-1c2d3e4f5a6b',
    'admin',
    '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj/VkklKiVSW',
    'Системный администратор',
    1
);

-- Оператор (пароль: operator123)
INSERT INTO users (id, username, password_hash, full_name, is_admin)
VALUES (
    'b7c8d9e0-1f2a-4b3c-8d4e-5f6a7b8c9d0e',
    'operator',
    '$2a$12$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uSp4R0rRq',
    'Оператор смены',
    0
);

-- Ключи шифрования MIFARE (демонстрационные)
INSERT INTO keys (id, name, key_value, key_type, sector, description)
VALUES
    ('c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f', 'Ключ-A-Сектор0', 'FFFFFFFFFFFF', 'A', 0, 'Стандартный ключ сектора 0'),
    ('d4e5f6a7-8b9c-4d0e-1f2a-3b4c5d6e7f8a', 'Ключ-A-Сектор1', 'A0A1A2A3A4A5', 'A', 1, 'Ключ баланса'),
    ('e7f8a9b0-1c2d-4e3f-4a5b-6c7d8e9f0a1b', 'Ключ-B-Сектор1', 'B0B1B2B3B4B5', 'B', 1, 'Ключ записи баланса');

-- Тестовые терминалы
INSERT INTO terminals (id, serial_number, name, location, route)
VALUES
    ('f0a1b2c3-4d5e-4f6a-7b8c-9d0e1f2a3b4c', 'TRM-001-BUS',  'Автобус №10 — Передний', 'Маршрут 10, депо Северное', '10'),
    ('a2b3c4d5-6e7f-4a8b-9c0d-1e2f3a4b5c6d', 'TRM-002-BUS',  'Автобус №10 — Задний',   'Маршрут 10, депо Северное', '10'),
    ('b5c6d7e8-9f0a-4b1c-2d3e-4f5a6b7c8d9e', 'TRM-003-TRAM', 'Трамвай №3',              'Ул. Ленина, остановка Центр', '3');

-- Тестовые карты
INSERT INTO cards (id, card_number, owner_name, balance, key_id)
VALUES
    ('c8d9e0f1-2a3b-4c4d-5e6f-7a8b9c0d1e2f', 'A1B2C3D4', 'Иванов Иван Иванович',    100000, 'c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f'),
    ('d1e2f3a4-5b6c-4d7e-8f9a-0b1c2d3e4f5a', 'E5F6A7B8', 'Петрова Мария Сергеевна',  50000, 'c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f'),
    ('e4f5a6b7-8c9d-4e0f-1a2b-3c4d5e6f7a8b', 'C9D0E1F2', 'Сидоров Алексей Петрович',     0, 'c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f');

-- Карта заблокированная
INSERT INTO cards (id, card_number, owner_name, balance, is_blocked, block_reason, key_id)
VALUES (
    'f7a8b9c0-1d2e-4f3a-4b5c-6d7e8f9a0b1c',
    'F3A4B5C6',
    'Козлов Дмитрий',
    2000,
    1,
    'Утеряна владельцем',
    'c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f'
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM cards;
DELETE FROM terminals;
DELETE FROM keys;
DELETE FROM users;
-- +goose StatementEnd
