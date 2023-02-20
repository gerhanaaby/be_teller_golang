package inits

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var Cfg Config

type Config struct {
	DBHost         string `env:"POSTGRES_HOST,required"`
	DBUserName     string `env:"POSTGRES_USER,required"`
	DBUserPassword string `env:"POSTGRES_PASSWORD,required"`
	DBName         string `env:"POSTGRES_DB,required"`
	DBPort         string `env:"POSTGRES_PORT,required"`
	ServerPort     string `env:"PORT,required"`

	ClientOrigin string `env:"CLIENT_ORIGIN"`

	TokenSecret    string        `env:"TOKEN_SECRET,required"`
	TokenExpiresIn time.Duration `env:"TOKEN_EXPIRED_IN,required"`
	TokenMaxAge    int           `env:"TOKEN_MAXAGE,required"`

	EmailFrom string `env:"EMAIL_FROM,required"`
	SMTPHost  string `env:"SMTP_HOST,required"`
	SMTPPass  string `env:"SMTP_PASS,required"`
	SMTPPort  int    `env:"SMTP_PORT,required"`
	SMTPUser  string `env:"SMTP_USER,required"`

	LogPath string
	EnvPath string
}

func LoadConfig(path string) (err error){

	if err := godotenv.Load(path); err != nil {
		fmt.Println(err)
		return err
	}

	err = env.Parse(&Cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigType("env")
// 	viper.SetConfigName("app")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }
