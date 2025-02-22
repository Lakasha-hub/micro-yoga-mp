package yoga

import (
	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/payment"
	"github.com/Lakasha-hub/micro-yoga-mp.git/pkg/mp"
)

type MembershipService struct {
	MercadoPagoClient *mp.Client
}

func NewMembershipService(mpClient *mp.Client) *MembershipService {
	return &MembershipService{
		MercadoPagoClient: mpClient,
	}
}

func (s *MembershipService) ProccessPayment(req payment.PaymentRequest) (payment.PaymentResponse, error) {
	pref, err := s.MercadoPagoClient.CreatePreference(req)
	if err != nil {
		return payment.PaymentResponse{
			Success: false,
			Message: "Error creating preference",
		}, err
	}

	return payment.PaymentResponse{
		Success: true,
		Message: "Preference created",
		Link:    pref.InitPoint,
	}, nil
}
