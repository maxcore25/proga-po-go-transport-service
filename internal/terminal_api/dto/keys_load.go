package dto

import "time"

// KeyLoadResponse — один ключ шифрования, передаваемый терминалу.
//
// Терминал использует эти ключи для расшифровки секторов карт MIFARE Classic.
// key_value передаётся в hex-кодировке (6 байт = 12 hex символов).
//
// @Description Single MIFARE encryption key for terminal
// @Name KeyLoadResponse
type KeyLoadResponse struct {
	// ID ключа в системе
	ID string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	// Человекочитаемое название
	Name string `json:"name" example:"Sector-1-Key-A"`
	// Hex-представление ключа (6 байт MIFARE Classic key)
	KeyValue string `json:"keyValue" example:"FFFFFFFFFFFF"`
	// Тип ключа: A (для чтения) или B (для записи)
	KeyType string `json:"keyType" example:"A"`
	// Номер сектора карты, к которому применяется ключ
	Sector int `json:"sector" example:"1"`
}

// KeysLoadResponse — полный пакет ключей для терминала.
//
// Терминал запрашивает этот список при старте и кэширует локально.
// Рекомендуется повторять запрос при каждом включении устройства.
//
// @Description Keys package delivered to terminal
// @Name KeysLoadResponse
type KeysLoadResponse struct {
	// Список активных ключей
	Keys []KeyLoadResponse `json:"keys"`
	// Количество ключей
	Count int `json:"count" example:"3"`
	// Время формирования пакета (UTC) — терминал должен логировать
	IssuedAt time.Time `json:"issuedAt"`
}
