package dto

// UpdateKeyRequest represents a key update request.
// @Description Key update request payload
// @Name UpdateKeyRequest
type UpdateKeyRequest struct {
	Name        *string `json:"name,omitempty" example:"Main Depot Key A"`
	KeyValue    *string `json:"keyValue,omitempty" example:"A1B2C3D4E5F6"`
	KeyType     *string `json:"keyType,omitempty" binding:"omitempty,oneof=A B" example:"A"`
	Sector      *int    `json:"sector,omitempty" example:"0"`
	Description *string `json:"description,omitempty" example:"Primary key for buses"`
	IsActive    *bool   `json:"isActive,omitempty" example:"true"`
}
