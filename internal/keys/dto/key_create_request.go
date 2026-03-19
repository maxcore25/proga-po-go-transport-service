package dto

// CreateKeyRequest represents a key creation request.
// @Description Key creation request payload
// @Name CreateKeyRequest
type CreateKeyRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=100" example:"Main Depot Key A"`
	KeyValue    string  `json:"keyValue" binding:"required,min=12,max=128" example:"A1B2C3D4E5F6"`
	KeyType     string  `json:"keyType" binding:"required,oneof=A B" example:"A"`
	Sector      int     `json:"sector" example:"0"`
	Description *string `json:"description,omitempty" example:"Primary key for buses"`
	IsActive    bool    `json:"isActive" example:"true"`
}
