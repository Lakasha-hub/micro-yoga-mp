package microyoga

import (
	"context"

	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga/mp"
	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga/payment"
)

type MembershipService struct {
	MercadoPagoClient *mp.Client
	repo              payment.PaymentRepo
}

func NewMembershipService(mpClient *mp.Client) *MembershipService {
	return &MembershipService{
		MercadoPagoClient: mpClient,
	}
}

func (s *MembershipService) ProccessPayment(ctx context.Context, req payment.PaymentRequest) (payment.PaymentResponse, error) {
	plan, err := s.MercadoPagoClient.CreateSuscriptionPlan(req)
	if err != nil {
		return payment.PaymentResponse{
			Success: false,
			Message: "Error creating plan subscription",
		}, err
	}
	s.repo.SavePayment(ctx, req)

	return payment.PaymentResponse{
		Success: true,
		Message: "Preference created",
		Link:    plan.InitPoint,
	}, nil
}
