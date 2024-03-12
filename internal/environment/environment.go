package environment

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type single struct {
	APPLICATION_PORT    string
	APPLICATION_ADDRESS string

	LOG_LEVEL string
	CORS_URL  string

	ENVIRONMENT string
	APP_VERSION string // nolint: golint

	MYSQL_DBNAME   string // nolint: golint
	MYSQL_HOST     string // nolint: golint
	MYSQL_PORT     int64  // nolint: golint
	MYSQL_USER     string // nolint: golint
	MYSQL_PASSWORD string // nolint: golint

	INTERVAL_GET_KEYS_TO_CREATE int64 // nolint: golint
}

func (e *single) Setup() {

	e.APPLICATION_PORT = getenv("APPLICATION_PORT", "8080")
	e.APPLICATION_ADDRESS = getenv("APPLICATION_ADDRESS", "localhost")

	e.LOG_LEVEL = getenv("LOG_LEVEL", "INFO")
	e.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	e.APP_VERSION = os.Getenv("APP_VERSION")

	e.MYSQL_DBNAME = getenv("MYSQL_DBNAME", "tempocerto_api")
	e.MYSQL_HOST = getenv("MYSQL_HOST", "localhost")
	e.MYSQL_PORT = getenvInt64("MYSQL_PORT", 3306)
	e.MYSQL_USER = getenv("MYSQL_USER", "root")
	e.MYSQL_PASSWORD = getenv("MYSQL_PASSWORD", "12345")

	e.INTERVAL_GET_KEYS_TO_CREATE = getenvInt64("INTERVAL_GET_KEYS_TO_CREATE", 10)

	e.CORS_URL = getenv("CORS_URL", "*")
}

func init() {
	fmt.Println(os.Getenv("ENVIRONMENT"))

	envVar := os.Getenv("ENVIRONMENT")

	if envVar == "development" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Println("Error loading .env.local file")
		}
	} else if envVar == "production" {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Println("Error loading .env.production file")
		}
	}
	env := GetInstance()
	env.Setup()
}

func (e *single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getenvInt64(key string, fallback int64) int64 {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	valueInt, _ := strconv.ParseInt(value, 10, 64)
	return valueInt
}

func GetenvFloat(key, fallback string) float64 {
	value := getenv(key, fallback)
	valueFloat, _ := strconv.ParseFloat(value, 32)
	return valueFloat
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created.")
		}
	}

	return singleInstance
}
