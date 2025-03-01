package payment

type (
	Payment struct {
		UserID            string  `json:"user_id"`
		Amount            float64 `json:"amount"`
		Currency          string  `json:"currency"`
		Status            string  `json:"status"`
		ExternalPaymentID string  `json:"external_payment_id"`
		CreatedAt         string  `json:"created_at"`
		UpdatedAt         string  `json:"updated_at"`
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
