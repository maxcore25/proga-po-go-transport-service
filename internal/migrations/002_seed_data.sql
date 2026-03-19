-- +goose Up
-- +goose StatementBegin

-- Администратор по умолчанию (пароль: admin123)
-- bcrypt hash сгенерирован с cost=12
INSERT INTO users (username, password_hash, full_name, is_admin)
VALUES (
    'admin',
    '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj/VkklKiVSW',
    'Системный администратор',
    1
);

-- Оператор (пароль: operator123)
INSERT INTO users (username, password_hash, full_name, is_admin)
VALUES (
    'operator',
    '$2a$12$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uSp4R0rRq',
    'Оператор смены',
    0
);

-- Ключи шифрования MIFARE (демонстрационные)
INSERT INTO keys (name, key_value, key_type, sector, description)
VALUES
    ('Ключ-A-Сектор0', 'FFFFFFFFFFFF', 'A', 0, 'Стандартный ключ сектора 0'),
    ('Ключ-A-Сектор1', 'A0A1A2A3A4A5', 'A', 1, 'Ключ баланса'),
    ('Ключ-B-Сектор1', 'B0B1B2B3B4B5', 'B', 1, 'Ключ записи баланса');

-- Тестовые терминалы
INSERT INTO terminals (serial_number, name, location, route)
VALUES
    ('TRM-001-BUS',  'Автобус №10 — Передний', 'Маршрут 10, депо Северное', '10'),
    ('TRM-002-BUS',  'Автобус №10 — Задний',   'Маршрут 10, депо Северное', '10'),
    ('TRM-003-TRAM', 'Трамвай №3',              'Ул. Ленина, остановка Центр', '3');

-- Тестовые карты
INSERT INTO cards (card_number, owner_name, balance, key_id)
VALUES
    ('A1B2C3D4', 'Иванов Иван Иванович',   100000, 1),  -- 1000.00 руб
    ('E5F6A7B8', 'Петрова Мария Сергеевна', 50000, 1),  -- 500.00 руб
    ('C9D0E1F2', 'Сидоров Алексей Петрович', 0,    1);  -- 0 руб (пустая)

-- Карта заблокированная
INSERT INTO cards (card_number, owner_name, balance, is_blocked, block_reason, key_id)
VALUES ('F3A4B5C6', 'Козлов Дмитрий',  2000, 1, 'Утеряна владельцем', 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM cards;
DELETE FROM terminals;
DELETE FROM keys;
DELETE FROM users;
-- +goose StatementEnd
