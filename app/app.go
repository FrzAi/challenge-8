package app

import (
	"challenge-8/config"
	"challenge-8/handler"
	"challenge-8/repository"
	"challenge-8/route"
	"challenge-8/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
