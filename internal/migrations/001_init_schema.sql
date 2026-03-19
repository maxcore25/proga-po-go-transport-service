-- +goose Up
-- +goose StatementBegin

-- =============================================
-- Таблица ключей шифрования MIFARE
-- Один ключ → много карт (1:N)
-- =============================================
CREATE TABLE IF NOT EXISTS keys (
    id          TEXT PRIMARY KEY,
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
    id            TEXT PRIMARY KEY,
    card_number   TEXT    NOT NULL UNIQUE,       -- UID карты (hex, например: A1B2C3D4)
    owner_name    TEXT    NOT NULL,              -- Имя владельца
    balance       INTEGER NOT NULL DEFAULT 0,   -- Баланс в копейках (избегаем float)
    is_blocked    INTEGER NOT NULL DEFAULT 0,   -- 0 = активна, 1 = заблокирована
    block_reason  TEXT,                         -- Причина блокировки
    key_id        TEXT NOT NULL REFERENCES keys(id) ON DELETE RESTRICT,
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
    id            TEXT PRIMARY KEY,
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
    id            TEXT PRIMARY KEY,
    card_id       TEXT NOT NULL REFERENCES cards(id)     ON DELETE RESTRICT,
    terminal_id   TEXT NOT NULL REFERENCES terminals(id) ON DELETE RESTRICT,
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
    id           TEXT PRIMARY KEY,
    username     TEXT    NOT NULL UNIQUE,
    password_hash TEXT   NOT NULL,              -- bcrypt hash
    full_name    TEXT    NOT NULL,
    is_admin     INTEGER NOT NULL DEFAULT 0,   -- 0 = обычный, 1 = администратор
    is_active    INTEGER NOT NULL DEFAULT 1,
    created_at   DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at   DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- =============================================
-- Таблица refresh_tokens
-- =============================================
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id         TEXT PRIMARY KEY,
    user_id    TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token      TEXT NOT NULL UNIQUE,
    expires_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_token   ON refresh_tokens(token);

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

DROP INDEX IF EXISTS idx_cards_card_number;
DROP INDEX IF EXISTS idx_cards_key_id;
DROP INDEX IF EXISTS idx_terminals_serial;
DROP INDEX IF EXISTS idx_transactions_card_id;
DROP INDEX IF EXISTS idx_transactions_terminal_id;
DROP INDEX IF EXISTS idx_transactions_created_at;
DROP INDEX IF EXISTS idx_refresh_tokens_user_id;
DROP INDEX IF EXISTS idx_refresh_tokens_token;

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS terminals;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS keys;
DROP TABLE IF EXISTS refresh_tokens;
-- +goose StatementEnd
