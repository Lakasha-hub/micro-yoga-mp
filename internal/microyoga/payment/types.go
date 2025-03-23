package payment

type (
	// Payment struct {
	// 	ID                uint      `json:"id"`
	// 	UserID            string    `json:"user_id"`
	// 	Amount            float64   `json:"amount"`
	// 	Currency          string    `json:"currency"`
	// 	Status            string    `json:"status"`
	// 	ExternalPaymentID string    `json:"external_payment_id"`
	// 	CreatedAt         time.Time `json:"created_at"`
	// 	UpdatedAt         time.Time `json:"updated_at"`
	// }

	PaymentRequest struct {
		UserEmail string  `json:"user_email" binding:"required"`
		Amount    float64 `json:"amount" binding:"required"`
		Currency  string  `json:"currency" binding:"required"`
	}

	Payment struct {
		UserID            string  `json:"user_id"`
		UserEmail         string  `json:"user_email"`
		Type              string  `json:"type"`
		ExternalPaymentID string  `json:"external_payment_id"`
		SubscriptionID    string  `json:"subscription_id"`
		Amount            float64 `json:"amount"`
		Currency          string  `json:"currency"`
		Status            string  `json:"status"`
	}
)
