package server

import (
	"log"
	"payment/controller"
	"payment/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Config *viper.Viper
	Router *gin.Engine
}

func InitHttpServer(config *viper.Viper, dbHandler *pgxpool.Conn) *HttpServer {

	serviceManager := services.NewServiceManager(dbHandler)
	controllerManager := controller.NewControllerManager(serviceManager)

	router := InitRouter(controllerManager)

	return &HttpServer{
		Config: config,
		Router: router,
	}
}

func (hs HttpServer) Start() {
	err := hs.Router.Run(hs.Config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
