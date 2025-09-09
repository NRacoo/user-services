package config

import (
	"os"
	"strconv"
	"user-service/common/util"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper/remote"
)

var Config AppConfig

type AppConfig struct {
	Port                  int      `json:"port"`
	AppName               string   `json:"appName"`
	AppEnv                string   `json:"appEnv"`
	SignatureKey          string   `json:"signatureKey"`
	Database              Database `json:"database"`
	RateLimiterMaxRequest float64  `json:"rateLimiterMaxRequest"`
	RateLimiterSecond     int      `json:"rateLimiterSecond"`
	JwtSecretKey          string   `json:"jwtSecretKey"`
	JwtExpirationTime     int      `json:"jwtExpirationTime"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnections    int    `json:"maxOpenConnections"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnections    int    `json:"maxIdleConnections"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func Init() { //panggil config.json
	_ = godotenv.Load()
	if os.Getenv("APP_PORT") != "" {
		Config = loadfromenv()
		return
	}
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind config")
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}

func loadfromenv() AppConfig {
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	rateLimiterMax, _ := strconv.ParseFloat(os.Getenv("RATE_LIMITER_MAX_REQUEST"), 64)
	rateLimiterSecond, _ := strconv.Atoi(os.Getenv("RATE_LIMITER_SECOND"))
	jwtexp, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))

	return AppConfig{
		Port:         port,
		AppName:      os.Getenv("APP_NAME"),
		AppEnv:       os.Getenv("APP_ENV"),
		SignatureKey: os.Getenv("SIGNATURE_KEY"),
		Database: Database{
			Host:                  os.Getenv("DB_HOST"),
			Port:                  dbPort,
			Name:                  os.Getenv("DB_NAME"),
			Username:              os.Getenv("DB_USER"),
			Password:              os.Getenv("DB_PASSWORD"),
			MaxOpenConnections:    10,
			MaxLifeTimeConnection: 10,
			MaxIdleConnections:    10,
			MaxIdleTime:           10,
		},
		RateLimiterMaxRequest: rateLimiterMax,
		RateLimiterSecond:     rateLimiterSecond,
		JwtSecretKey:          os.Getenv("JWT_SECRET_KEY"),
		JwtExpirationTime:     jwtexp,
	}
}
