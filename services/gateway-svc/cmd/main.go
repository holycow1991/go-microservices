package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/r0tten911/gateway-svc/pkg/auth"
	"github.com/r0tten911/gateway-svc/pkg/config"
	"github.com/r0tten911/gateway-svc/pkg/order"
	"github.com/r0tten911/gateway-svc/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
