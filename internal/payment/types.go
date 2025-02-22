package payment

import "gorm.io/gorm"

type (
	Payment struct {
		gorm.Model
		UserID     string  `json:"user_id" gorm:"not null"`
		Amount     float64 `json:"amount" gorm:"not null"`
		Currency   string  `json:"currency" gorm:"not null"`
		Status     string  `json:"status" gorm:"not null" default:"pending"`
		ExternalID string  `json:"external_payment_id" gorm:"not null"`
	}

	PaymentRequest struct {
		UserID   string  `json:"user_id" binding:"required"`
		Amount   float64 `json:"amount" binding:"required"`
		Currency string  `json:"currency" binding:"required"`
	}

	PaymentResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Link    string `json:"link"`
	}
)
