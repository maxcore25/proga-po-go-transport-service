package dto

// PaymentAuthRequest — запрос от физического терминала на авторизацию платежа.
//
// Терминал уже расшифровал карту локально (используя ключи из /keys),
// получил card_number (UID карты) и отправляет его вместе с суммой.
//
// @Description Terminal payment authorization request
// @Name PaymentAuthRequest
type PaymentAuthRequest struct {
	// UID карты MIFARE (hex, например: A1B2C3D4)
	CardNumber string `json:"cardNumber" binding:"required,min=4,max=32" example:"A1B2C3D4"`
	// Сумма списания в копейках (> 0)
	Amount int `json:"amount" binding:"required,gt=0" example:"3500"`
	// Серийный номер терминала — идентифицирует устройство
	TerminalSerial string `json:"terminalSerial" binding:"required" example:"TRM-001-BUS"`
}
