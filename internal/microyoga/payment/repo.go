package payment

import (
	"context"
	"database/sql"
)

type PaymentRepo interface {
	Save(ctx context.Context, payment Payment) error
}

type PaymentMysqlRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) PaymentRepo {
	return &PaymentMysqlRepo{db: db}
}

func (r *PaymentMysqlRepo) Save(ctx context.Context, payment Payment) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO payments (user_id, amount, currency, status, external_payment_id, user_email) VALUES (?, ?, ?, ?, ?)",
		payment.UserID,
		payment.Amount,
		payment.Currency,
		payment.Status,
		payment.ExternalPaymentID,
		payment.UserEmail,
	)

	if err != nil {
		return err
	}
	return nil
}
