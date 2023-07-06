package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `json:"APP_ENV"`
	ServerAddress          string `json:"SERVER_ADDRESS"`
	ContextTimeout         int    `json:"CONTEXT_TIMEOUT"`
	DBHost                 string `json:"DB_HOST"`
	DBPort                 string `json:"DB_PORT"`
	DBUser                 string `json:"DB_USER"`
	DBPass                 string `json:"DB_PASS"`
	DBName                 string `json:"DB_NAME"`
	AccessTokenExpiryHour  int    `json:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `json:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `json:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `json:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Oops! can't find the specified env file.", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Failed while loading the environment.", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env.")
	}

	if env.AppEnv == "production" {
		log.Println("The App is running in production env.")
	}

	return &env
}
