package inits

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-20
 * @desc [Model Config, LoadEnv dan Parse Env ke Config]
 */
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

	LogPerformancePath string
	LogReportPath      string
	Base64Path         string
	EnvPath            string
}

func LoadConfig(path string) (err error) {

	if err := godotenv.Load(path); err != nil {
		return err
	}

	err = env.Parse(&Cfg)
	if err != nil {
		return err
	}

	return nil
}
