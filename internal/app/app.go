package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wartinich/book-archive/docs"
	"github.com/wartinich/book-archive/internal/delivery/http"
)

// @title           Book Archive API
// @version         1.0
// @description     Book Archive API
// @host            127.0.0.1:8000
// @BasePath  		/
func Run() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error init config: %s", err.Error())
	}

	r := http.Handler()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	go func() {}()
	r.Run(viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
