package main

import (
	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/config"
	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/handlers"
	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/platform/db"
	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/yoga"
	"github.com/Lakasha-hub/micro-yoga-mp.git/pkg/mp"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db.Init(cfg)

	// Config MercadoPago
	mpClient, err := mp.NewClient(cfg.MPAccessToken)
	if err != nil {
		panic("Error creating MercadoPago client")
	}

	// Display services and handlers
	membershipService := yoga.NewMembershipService(mpClient)
	paymentHandler := handlers.NewPaymentHandler(membershipService)

	// Create a new Gin router
	router := gin.Default()

	// Define the route
	router.POST("/payment", paymentHandler.ProccessPayment)

	// Initialize the server
	router.Run(":%s", cfg.ServerPort)
}
