package mp

import (
	"context"

	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga/payment"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

type Client struct {
	preapprovalClient preapproval.Client
}

func NewClient(accessToken string) (*Client, error) {
	cfg, err := config.New(accessToken)
	if err != nil {
		return nil, err
	}

	return &Client{
		preapprovalClient: preapproval.NewClient(cfg),
	}, nil
}

func (c *Client) CreateSuscriptionPlan(req payment.PaymentRequest) (*preapproval.Response, error) {
	planRequest := preapproval.Request{
		Reason: "Formación Natha Yoga",
		AutoRecurring: &preapproval.AutoRecurringRequest{
			Frequency:         1,
			FrequencyType:     "months",
			TransactionAmount: req.Amount,
			CurrencyID:        req.Currency,
		},
	}

	res, err := c.preapprovalClient.Create(context.Background(), planRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}
