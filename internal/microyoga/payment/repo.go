package payment

import (
	"context"
	"database/sql"
)

type PaymentRepo interface {
	SavePayment(ctx context.Context, payment PaymentRequest) error
}

type PaymentMysqlRepo struct {
	db *sql.DB
}

func NewPaymentMysqlRepo(db *sql.DB) PaymentRepo {
	return &PaymentMysqlRepo{db: db}
}

func (r *PaymentMysqlRepo) SavePayment(ctx context.Context, payment PaymentRequest) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO payments (user_id, amount, currency, status) VALUES (?, ?, ?)",
		payment.UserID,
		payment.Amount,
		payment.Currency,
	)
	if err != nil {
		return err
	}
	return nil
}
