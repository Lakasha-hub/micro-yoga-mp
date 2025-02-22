package handlers

import (
	"net/http"

	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/payment"
	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/yoga"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	MembershipService *yoga.MembershipService
}

func NewPaymentHandler(ms *yoga.MembershipService) *PaymentHandler {
	return &PaymentHandler{
		MembershipService: ms,
	}
}

func (h *PaymentHandler) ProccessPayment(ctx *gin.Context) {
	var req payment.PaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	res, err := h.MembershipService.ProccessPayment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payload": res})
}
