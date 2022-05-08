package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wartinich/web/docs"
	"github.com/wartinich/web/pkg/handler"
)

// @title           Book Archive API
// @version         1.0
// @description     Book Archive API
// @host            127.0.0.1:8000
// @BasePath  		/
func main() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error init config: %s", err.Error())
	}

	r := handler.Handler()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
