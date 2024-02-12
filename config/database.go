package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"restful-api-golang/helper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
}

func ConnectionDB() *gorm.DB {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	configModel := Config{}
	err = config.Unmarshal(&configModel)
	helper.PanicIfError(err)

	sql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configModel.DBUser, configModel.DBPassword,
		configModel.DBHost, configModel.DBPort, configModel.DBName)

	db, err := gorm.Open(mysql.Open(sql), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}
