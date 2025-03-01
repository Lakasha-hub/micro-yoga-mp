package handlers

import (
	"net/http"

	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga"
	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga/payment"
	"github.com/gin-gonic/gin"
)

func NewProccessPayment(useCase microyoga.MembershipService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req payment.PaymentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
			return
		}

		res, err := useCase.ProccessPayment(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing payment"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"payload": res})
	}

}
