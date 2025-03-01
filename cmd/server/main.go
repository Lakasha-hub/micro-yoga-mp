package main

import (
	"fmt"
	"log"

	"github.com/Lakasha-hub/micro-yoga-mp/internal/handlers"
	yoga "github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga"
	"github.com/Lakasha-hub/micro-yoga-mp/internal/platform/db"
	"github.com/Lakasha-hub/micro-yoga-mp/pkg/mp"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() {

	cfg := LoadConfig()

	db, err := db.NewDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		panic("Error connecting to the database")
	}

	mpClient, err := mp.NewClient(cfg.MPAccessToken)
	if err != nil {
		panic("Error creating MercadoPago client")
	}

	membershipService := yoga.NewMembershipService(mpClient)
	paymentHandler := handlers.NewPaymentHandler(membershipService)

	router := gin.Default()

	router.POST("/payment", paymentHandler.ProccessPayment)

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
