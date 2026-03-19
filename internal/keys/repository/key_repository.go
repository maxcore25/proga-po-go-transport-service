package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/model"
	"gorm.io/gorm"
)

type KeyRepository interface {
	Create(key *model.Key) error
	GetByID(id uuid.UUID) (*model.Key, error)
	GetByKeyValue(keyValue string) (*model.Key, error)
	GetAll() ([]*model.Key, error)
	// GetAllActive возвращает только активные ключи — для отправки на терминалы.
	GetAllActive() ([]*model.Key, error)
	Find(filter dto.KeyFilter) ([]*model.Key, error)
	UpdateByID(id uuid.UUID, updateData dto.UpdateKeyRequest) error
	DeleteByID(id uuid.UUID) error
}

type keyRepository struct {
	db *gorm.DB
}

func NewKeyRepository(db *gorm.DB) KeyRepository {
	return &keyRepository{db: db}
}

func (r *keyRepository) Create(key *model.Key) error {
	return r.db.Create(key).Error
}

func (r *keyRepository) GetByID(id uuid.UUID) (*model.Key, error) {
	var k model.Key
	if err := r.db.First(&k, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *keyRepository) GetByKeyValue(keyValue string) (*model.Key, error) {
	var k model.Key
	if err := r.db.First(&k, "key_value = ?", keyValue).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *keyRepository) GetAll() ([]*model.Key, error) {
	var keys []*model.Key
	if err := r.db.Find(&keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

// GetAllActive фильтрует только is_active=true.
// Отозванные ключи (is_active=false) не передаются терминалам —
// карты на таких ключах не смогут быть расшифрованы (защита от компрометации).
func (r *keyRepository) GetAllActive() ([]*model.Key, error) {
	var keys []*model.Key
	if err := r.db.Where("is_active = ?", true).Find(&keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *keyRepository) Find(filter dto.KeyFilter) ([]*model.Key, error) {
	db := r.db.Model(&model.Key{})

	if filter.IsActive != nil {
		db = db.Where("is_active = ?", *filter.IsActive)
	}
	if filter.KeyType != nil {
		db = db.Where("key_type = ?", *filter.KeyType)
	}

	var keys []*model.Key
	if err := db.Find(&keys).Error; err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *keyRepository) UpdateByID(id uuid.UUID, updateData dto.UpdateKeyRequest) error {
	return r.db.Model(&model.Key{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *keyRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Key{}).Error
}
