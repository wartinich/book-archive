package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wartinich/book-archive/internal/delivery/http"
)

func Run() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error init config: %s", err.Error())
	}

	r := http.Handler()

	go func() {}()
	r.Run(viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
