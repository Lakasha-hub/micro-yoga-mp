package payment

import (
	"gorm.io/gorm"
)

type PaymentRepo interface {
	SavePayment(payment Payment) error
}

type PaymentMysqlRepo struct {
	db *gorm.DB
}

func NewPaymentMysqlRepo(db *gorm.DB) PaymentRepo {
	return &PaymentMysqlRepo{db: db}
}

func (r *PaymentMysqlRepo) SavePayment(payment Payment) error {
	result := r.db.Create(&payment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
