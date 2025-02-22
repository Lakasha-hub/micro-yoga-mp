package mp

import (
	"context"

	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/payment"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

type Client struct {
	preferenceClient preference.Client
}

func NewClient(accessToken string) (*Client, error) {
	cfg, err := config.New(accessToken)
	if err != nil {
		return nil, err
	}

	return &Client{
		preferenceClient: preference.NewClient(cfg),
	}, nil
}

func (c *Client) CreatePreference(req payment.PaymentRequest) (*preference.Response, error) {
	prefRequest := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:      "Yoga Membership",
				Quantity:   1,
				UnitPrice:  req.Amount,
				CurrencyID: req.Currency,
			},
		},
		BackURLs: &preference.BackURLsRequest{
			Success: "http://localhost:8080/success",
			Pending: "http://localhost:8080/pending",
			Failure: "http://localhost:8080/failure",
		},
		AutoReturn: "approved",
	}

	pref, err := c.preferenceClient.Create(context.Background(), prefRequest)
	if err != nil {
		return nil, err
	}

	return pref, nil
}
