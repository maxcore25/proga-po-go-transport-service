-- +goose Up
-- +goose StatementBegin

-- =============================================
-- Таблица ключей шифрования MIFARE
-- Один ключ → много карт (1:N)
-- =============================================
CREATE TABLE IF NOT EXISTS keys (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT    NOT NULL,               -- Человекочитаемое название ключа
    key_value   TEXT    NOT NULL,               -- Hex-представление ключа (48 hex chars = 6 bytes MIFARE key)
    key_type    TEXT    NOT NULL DEFAULT 'A',   -- 'A' или 'B' (MIFARE Classic sector keys)
    sector      INTEGER NOT NULL DEFAULT 0,     -- Номер сектора карты
    description TEXT,
    is_active   INTEGER NOT NULL DEFAULT 1,     -- 1 = активен, 0 = отозван
    created_at  DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at  DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- =============================================
-- Таблица транспортных карт MIFARE
-- =============================================
CREATE TABLE IF NOT EXISTS cards (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    card_number   TEXT    NOT NULL UNIQUE,       -- UID карты (hex, например: A1B2C3D4)
    owner_name    TEXT    NOT NULL,              -- Имя владельца
    balance       INTEGER NOT NULL DEFAULT 0,   -- Баланс в копейках (избегаем float)
    is_blocked    INTEGER NOT NULL DEFAULT 0,   -- 0 = активна, 1 = заблокирована
    block_reason  TEXT,                         -- Причина блокировки
    key_id        INTEGER NOT NULL REFERENCES keys(id) ON DELETE RESTRICT,
    expires_at    DATETIME,                     -- Срок действия карты
    created_at    DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at    DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE INDEX IF NOT EXISTS idx_cards_card_number ON cards(card_number);
CREATE INDEX IF NOT EXISTS idx_cards_key_id      ON cards(key_id);

-- =============================================
-- Таблица терминалов (валидаторов)
-- =============================================
CREATE TABLE IF NOT EXISTS terminals (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    serial_number TEXT    NOT NULL UNIQUE,      -- Серийный номер устройства
    name          TEXT    NOT NULL,             -- Название (например: "Автобус №42 - Передний")
    location      TEXT    NOT NULL,             -- Адрес установки
    route         TEXT,                         -- Маршрут/линия
    is_active     INTEGER NOT NULL DEFAULT 1,   -- 0 = деактивирован
    last_seen_at  DATETIME,                     -- Последнее обращение к серверу
    created_at    DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at    DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE INDEX IF NOT EXISTS idx_terminals_serial ON terminals(serial_number);

-- =============================================
-- Таблица транзакций
-- =============================================
CREATE TABLE IF NOT EXISTS transactions (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    card_id       INTEGER NOT NULL REFERENCES cards(id)     ON DELETE RESTRICT,
    terminal_id   INTEGER NOT NULL REFERENCES terminals(id) ON DELETE RESTRICT,
    amount        INTEGER NOT NULL,             -- Сумма списания в копейках
    balance_before INTEGER NOT NULL,            -- Баланс до транзакции
    balance_after  INTEGER NOT NULL,            -- Баланс после транзакции
    status        TEXT    NOT NULL DEFAULT 'approved', -- 'approved' | 'declined'
    decline_reason TEXT,                        -- Причина отказа (если status = 'declined')
    created_at    DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE INDEX IF NOT EXISTS idx_transactions_card_id     ON transactions(card_id);
CREATE INDEX IF NOT EXISTS idx_transactions_terminal_id ON transactions(terminal_id);
CREATE INDEX IF NOT EXISTS idx_transactions_created_at  ON transactions(created_at);

-- =============================================
-- Таблица пользователей системы
-- =============================================
CREATE TABLE IF NOT EXISTS users (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    username     TEXT    NOT NULL UNIQUE,
    password_hash TEXT   NOT NULL,              -- bcrypt hash
    full_name    TEXT    NOT NULL,
    is_admin     INTEGER NOT NULL DEFAULT 0,   -- 0 = обычный, 1 = администратор
    is_active    INTEGER NOT NULL DEFAULT 1,
    created_at   DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at   DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- =============================================
-- Триггеры: автоматически обновляем updated_at
-- =============================================
CREATE TRIGGER IF NOT EXISTS trg_cards_updated_at
    AFTER UPDATE ON cards
    FOR EACH ROW
BEGIN
    UPDATE cards SET updated_at = datetime('now') WHERE id = NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_terminals_updated_at
    AFTER UPDATE ON terminals
    FOR EACH ROW
BEGIN
    UPDATE terminals SET updated_at = datetime('now') WHERE id = NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_keys_updated_at
    AFTER UPDATE ON keys
    FOR EACH ROW
BEGIN
    UPDATE keys SET updated_at = datetime('now') WHERE id = NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_users_updated_at
    AFTER UPDATE ON users
    FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = datetime('now') WHERE id = NEW.id;
END;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_users_updated_at;
DROP TRIGGER IF EXISTS trg_keys_updated_at;
DROP TRIGGER IF EXISTS trg_terminals_updated_at;
DROP TRIGGER IF EXISTS trg_cards_updated_at;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS terminals;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS keys;
-- +goose StatementEnd
