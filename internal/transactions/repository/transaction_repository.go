package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetByID(id uuid.UUID) (*model.Transaction, error)
	GetAll() ([]*model.Transaction, error)
	Find(filter dto.TransactionFilter) ([]*model.Transaction, error)
	UpdateByID(id uuid.UUID, updateData dto.UpdateTransactionRequest) error
	DeleteByID(id uuid.UUID) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetByID(id uuid.UUID) (*model.Transaction, error) {
	var t model.Transaction
	if err := r.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *transactionRepository) GetAll() ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) Find(filter dto.TransactionFilter) ([]*model.Transaction, error) {
	db := r.db.Model(&model.Transaction{})

	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}

	var transactions []*model.Transaction
	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) UpdateByID(id uuid.UUID, updateData dto.UpdateTransactionRequest) error {
	updates := map[string]interface{}{}

	if updateData.CardID != nil {
		cardID, err := uuid.Parse(*updateData.CardID)
		if err != nil {
			return err
		}
		updates["card_id"] = cardID
	}
	if updateData.TerminalID != nil {
		terminalID, err := uuid.Parse(*updateData.TerminalID)
		if err != nil {
			return err
		}
		updates["terminal_id"] = terminalID
	}
	if updateData.Amount != nil {
		updates["amount"] = *updateData.Amount
	}
	if updateData.BalanceBefore != nil {
		updates["balance_before"] = *updateData.BalanceBefore
	}
	if updateData.BalanceAfter != nil {
		updates["balance_after"] = *updateData.BalanceAfter
	}
	if updateData.Status != nil {
		updates["status"] = *updateData.Status
	}
	if updateData.DeclineReason != nil {
		updates["decline_reason"] = *updateData.DeclineReason
	}

	return r.db.Model(&model.Transaction{}).Where("id = ?", id).Updates(updates).Error
}

func (r *transactionRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Transaction{}).Error
}
